package telecommands

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hungdoo/bot/src/common"
	"github.com/hungdoo/bot/src/packages/command/balance"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/command/contract"
	"github.com/hungdoo/bot/src/packages/command/debank"
	"github.com/hungdoo/bot/src/packages/command/tomb"
	"github.com/hungdoo/bot/src/packages/db"
	"github.com/hungdoo/bot/src/packages/log"
	"go.mongodb.org/mongo-driver/bson"
)

type CommandMap map[string]command.ICommand

func (cm *CommandMap) Search(nameSubString string) []command.ICommand {
	list := make([]command.ICommand, 0)
	for _, v := range *cm {
		if strings.Contains(v.GetName(), nameSubString) {
			list = append(list, v)
		}
	}
	return list
}
func (cm *CommandMap) ToList() []command.ICommand {
	list := make([]command.ICommand, 0)
	for _, v := range *cm {
		list = append(list, v)
	}
	return list
}
func (cm *CommandMap) Filter(cmdType command.CommandType) *CommandMap {
	filtered := CommandMap{}
	for _, v := range *cm {
		if v.GetType() == cmdType {
			filtered[v.GetName()] = v
		}
	}
	return &filtered
}
func (cm *CommandMap) ToActiveList() []command.ICommand {
	actives := make([]command.ICommand, 0)
	for _, v := range *cm {
		if v.IsEnabled() {
			actives = append(actives, v)
		}
	}
	return actives
}

// @dev internal facing controller
// Responsibility: validates, processes commands & handles errors
type CommandFactory struct {
	commands          CommandMap
	lastRefreshedAt   time.Time
	refreshDbInterval time.Duration
}

func NewCommadFactory() CommandFactory {
	return CommandFactory{commands: map[string]command.ICommand{}, lastRefreshedAt: time.Now(), refreshDbInterval: 3 * time.Hour}
}

func (c *CommandFactory) Add(cmdType command.CommandType, messages []string) string {
	if len(messages) < 2 {
		return "Add needs at least 2 params"
	}
	name := messages[0]
	messages = messages[1:]
	if v, ok := c.commands[name]; ok {
		if err := v.SetData(messages); err != nil {
			return err.Error()
		}
		filter := bson.M{"name": v.GetName()}
		update := bson.M{"$set": bson.M{"data": v.GetData()}}
		if err := db.GetDb().Update("commands", filter, update); err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Command [%v] updated", name)
	} else {
		var newCommand command.ICommand

		switch cmdType {
		case command.ContractCall:
			newCommand = &contract.Command{
				Command: command.Command{
					Name:     name,
					Enabled:  true,
					Type:     cmdType,
					IdleTime: time.Second * 60,
				},
			}
		case command.Tomb:
			newCommand = &tomb.TombCommand{
				Command: command.Command{
					Name:     name,
					Enabled:  true,
					Type:     cmdType,
					IdleTime: time.Minute * 5,
				},
			}
		case command.Debank:
			newCommand = &debank.Command{
				Command: command.Command{
					Name:     name,
					Enabled:  true,
					Type:     cmdType,
					IdleTime: time.Second * 60,
				},
			}
		case command.Balance:
			newCommand = &balance.BalanceCommand{
				Command: command.Command{
					Name:     name,
					Enabled:  true,
					Type:     cmdType,
					IdleTime: time.Second * 30,
				},
			}
		}
		if newCommand == nil {
			return fmt.Sprintf("Command [%v] failed to add", name)
		}

		if err := newCommand.SetData(messages); err != nil {
			return err.Error()
		}
		if err := db.GetDb().Insert("commands", newCommand.GetUnderlying()); err != nil {
			return err.Error()
		}
		c.commands[name] = newCommand
		return fmt.Sprintf("Command [%v] added", name)
	}
}

func (c *CommandFactory) Remove(name string) string {
	if _, ok := c.commands[name]; ok {
		deleteQuery := bson.M{"name": name}
		if err := db.GetDb().Delete("commands", deleteQuery); err != nil {
			return err.Error()
		}
		delete(c.commands, name)
		return fmt.Sprintf("Command [%v] removed", name)
	}
	return fmt.Sprintf("Command [%v] not found", name)
}

func (c *CommandFactory) Show(name string) string {
	searchedList := c.commands.Search(name)
	if len(searchedList) != 0 {
		b, err := json.MarshalIndent(searchedList, "", "  ")
		if err != nil {
			return err.Error()
		}
		return string(b)
	}
	return fmt.Sprintf("Command [%v] not found", name)
}

func (c *CommandFactory) Exec(cmdType command.CommandType, task string, opts ...string) (res string, err error) {
	filtered := c.commands.Filter(cmdType)
	switch cmdType {
	case command.Tomb:
		subCmd := ""
		if len(opts) != 0 {
			subCmd = opts[0]
		}
		searchedList := filtered.Search(task)

		var executed []string
		for _, cmd := range searchedList {
			result, execErr := cmd.Execute(true, subCmd)
			if execErr != nil && execErr.Level >= common.Error {
				log.GeneralLogger.Printf("Job [%s] exec failed: [%s]", cmd.GetName(), execErr.Error())
				if execErr.Level >= common.Critical {
					executed = append(executed, fmt.Sprintf("%v with reason %s", c.Off(cmd.GetName()), execErr.Error()))
				} else {
					executed = append(executed, execErr.Error())
				}
				continue
			}
			if result != "" {
				cmd.SetDisplayMsg(result)
				executed = append(executed, fmt.Sprintf("job [%s] result: %s", cmd.GetName(), result))
			}
		}

		return string(strings.Join(executed, "\n")), nil
	default:
		searchedList := filtered.Search(task)

		var executed []string
		for _, cmd := range searchedList {
			result, execErr := cmd.Execute(true, "")
			if execErr != nil && execErr.Level >= common.Error {
				log.GeneralLogger.Printf("Job [%s] exec failed: [%s]", cmd.GetName(), execErr.Error())
				if execErr.Level >= common.Critical {
					executed = append(executed, fmt.Sprintf("%v with reason %s", c.Off(cmd.GetName()), execErr.Error()))
				} else {
					executed = append(executed, execErr.Error())
				}
				continue
			}
			if result != "" {
				cmd.SetDisplayMsg(result)
				executed = append(executed, fmt.Sprintf("job [%s] result: %s", cmd.GetName(), result))
			}
		}

		return string(strings.Join(executed, "\n")), nil
	}
}

func (c *CommandFactory) List() string {
	enabledCMD := []string{}
	disabledCMD := []string{}
	for _, cmd := range c.commands {
		if cmd.IsEnabled() {
			enabledCMD = append(enabledCMD, fmt.Sprintf("[+] %v", cmd.GetOverview()))
		} else {
			disabledCMD = append(disabledCMD, fmt.Sprintf("[-] %v", cmd.GetOverview()))
		}
	}
	return strings.Join(append(append(enabledCMD, ""), disabledCMD...), "\n")
}

func (c *CommandFactory) On(name string) string {
	if v, ok := c.commands[name]; ok {
		v.SetEnabled(true)
		filter := bson.M{"name": v.GetName()}
		update := bson.M{"$set": bson.M{"enabled": true}}
		if err := db.GetDb().Update("commands", filter, update); err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Command [%v] on", name)
	}
	return fmt.Sprintf("Command [%v] not found", name)
}

func (c *CommandFactory) Off(name string) string {
	if v, ok := c.commands[name]; ok {
		v.SetEnabled(false)
		filter := bson.M{"name": v.GetName()}
		update := bson.M{"$set": bson.M{"enabled": false}}
		if err := db.GetDb().Update("commands", filter, update); err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Command [%v] off", name)
	}
	return fmt.Sprintf("Command [%v] not found", name)
}

func (c *CommandFactory) SetInterval(name string, interval time.Duration) string {
	if v, ok := c.commands[name]; ok {
		v.SetIdleTime(interval)
		filter := bson.M{"name": v.GetName()}
		update := bson.M{"$set": bson.M{"idletime": interval}}
		if err := db.GetDb().Update("commands", filter, update); err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Command [%v] interval: [%v]", name, interval)
	}
	return fmt.Sprintf("Command [%v] not found", name)
}

func (c *CommandFactory) GetJobs() ([]command.ICommand, error) {
	if len(c.commands) == 0 || time.Since(c.lastRefreshedAt) >= c.refreshDbInterval {
		c.lastRefreshedAt = time.Now()

		cursor, err := db.GetDb().Find("commands", bson.D{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {
			cmd := command.Command{}
			if err := cursor.Decode(&cmd); err != nil {
				return nil, err
			}
			var _command command.ICommand
			switch cmd.Type {
			case command.ContractCall:
				_command = &contract.Command{
					Command: cmd,
				}
			case command.Tomb:
				_command = &tomb.TombCommand{
					Command: cmd,
				}
				_command.SetData(_command.GetData())
			case command.Balance:
				_command = &balance.BalanceCommand{
					Command: cmd,
				}
				_command.SetData(_command.GetData())
			case command.Debank:
				_command = &debank.Command{
					Command: cmd,
				}
			}
			if _command == nil {
				continue
			}
			name := cmd.GetName()
			c.commands[name] = _command
			log.GeneralLogger.Printf("Loaded Command [%v]\n", name)
		}
	}

	return c.commands.ToActiveList(), nil
}
