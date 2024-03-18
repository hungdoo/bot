package telecommands

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hungdoo/bot/src/common"
	"github.com/hungdoo/bot/src/packages/command/balance"
	"github.com/hungdoo/bot/src/packages/command/bybitido"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/command/contract"
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
	commands CommandMap
}

func NewCommandFactory() CommandFactory {
	return CommandFactory{commands: map[string]command.ICommand{}}
}

func (c *CommandFactory) Add(cmdType command.CommandType, messages []string) string {
	if len(messages) < 1 {
		return "Add needs at least 1 params"
	}
	name := messages[0]
	var data []string
	if len(messages) > 1 {
		data = messages[1:]
	}
	if v, ok := c.commands[name]; ok {
		if err := v.SetData(data); err != nil {
			return err.Error()
		}
		if err := UpdateCmd(v); err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Command [%v] updated", name)
	} else {
		var newCommand command.ICommand

		switch cmdType {
		case command.ContractCall:
			newCommand = &contract.ContractCommand{
				Id: name,
				Command: command.Command{
					Name:     name,
					Enabled:  true,
					Type:     cmdType,
					IdleTime: time.Second * 60,
				},
			}
		case command.Tomb:
			newCommand = &tomb.TombCommand{
				Id: name,
				Command: command.Command{
					Name:     name,
					Enabled:  true,
					Type:     cmdType,
					IdleTime: time.Minute * 1,
				},
			}
		case command.Balance:
			newCommand = &balance.BalanceCommand{
				Id: name,
				Command: command.Command{
					Name:     name,
					Enabled:  true,
					Type:     cmdType,
					IdleTime: time.Second * 30,
				},
			}
		case command.BybitIdo:
			newCommand = &bybitido.IdoCommand{
				Id: name,
				Command: command.Command{
					Name:     name,
					Enabled:  true,
					Type:     cmdType,
					IdleTime: time.Hour * 1,
				},
			}
		}
		if newCommand == nil {
			return fmt.Sprintf("Command [%v] failed to add", name)
		}

		if err := newCommand.SetData(data); err != nil {
			return err.Error()
		}

		if err := StoreCmd(newCommand); err != nil {
			return err.Error()
		}
		c.commands[name] = newCommand
		return fmt.Sprintf("Command [%v] added", name)
	}
}

func (c *CommandFactory) Remove(name string) string {
	if _, ok := c.commands[name]; ok {
		deleteQuery := bson.M{"_id": name}
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
	var executedCmds []command.ICommand
	var executedResults []string

	switch cmdType {
	case command.Tomb, command.BybitIdo:
		subCmd := ""
		if len(opts) != 0 {
			subCmd = opts[0]
		}
		searchedList := filtered.Search(task)
		if len(searchedList) == 0 {
			return fmt.Sprintf("Task [%v] not found", task), nil
		}

		for _, cmd := range searchedList {
			result, execErr := cmd.Execute(true, subCmd)
			if execErr != nil && execErr.Level >= common.Error {
				log.GeneralLogger.Printf("Job [%s] exec failed: [%s]", cmd.GetName(), execErr.Error())
				if execErr.Level >= common.Critical {
					executedResults = append(executedResults, fmt.Sprintf("%v with reason %s", c.Off(cmd.GetName()), execErr.Error()))
				} else {
					executedResults = append(executedResults, fmt.Sprintf("%v failed with [%s:%s]", cmd.GetName(), execErr.Level, execErr.Error()))
				}
				continue
			}

			// exec seccessfully -> update db
			if result != "" {
				cmd.SetDisplayMsg(result)
				executedResults = append(executedResults, fmt.Sprintf("[%s]\n%s", cmd.GetName(), result))
			}
			executedCmds = append(executedCmds, cmd)
		}

	default:
		searchedList := filtered.Search(task)
		if len(searchedList) == 0 {
			return fmt.Sprintf("Task [%v] not found", task), nil
		}

		for _, cmd := range searchedList {
			result, execErr := cmd.Execute(true, "")
			if execErr != nil && execErr.Level >= common.Error {
				log.GeneralLogger.Printf("Job [%s] exec failed: [%s]", cmd.GetName(), execErr.Error())
				if execErr.Level >= common.Critical {
					executedResults = append(executedResults, fmt.Sprintf("%v with reason %s", c.Off(cmd.GetName()), execErr.Error()))
				} else {
					executedResults = append(executedResults, execErr.Error())
				}
				continue
			}

			// exec seccessfully -> update db
			if result != "" {
				cmd.SetDisplayMsg(result)
				executedResults = append(executedResults, fmt.Sprintf("[%s]\n%s", cmd.GetName(), result))
			}
			executedCmds = append(executedCmds, cmd)
		}
	}
	if len(executedCmds) != 0 {
		UpdateMultiCmd(executedCmds)
	}
	return string(strings.Join(executedResults, "\n")), nil
}

func (c *CommandFactory) List(showAll bool) string {
	enabledCMD := []string{}
	disabledCMD := []string{}
	for _, cmd := range c.commands {
		if cmd.IsEnabled() {
			enabledCMD = append(enabledCMD, fmt.Sprintf("[+] %v", cmd.GetOverview()))
		} else {
			disabledCMD = append(disabledCMD, fmt.Sprintf("[-] %v", cmd.GetOverview()))
		}
	}

	if showAll {
		return strings.Join(append(append(enabledCMD, ""), disabledCMD...), "\n")
	} else {
		return strings.Join(enabledCMD, "\n")
	}
}

func (c *CommandFactory) On(name string) string {
	if v, ok := c.commands[name]; ok {
		v.SetEnabled(true)
		if err := UpdateCmd(v); err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Command [%v] on", name)
	}
	return fmt.Sprintf("Command [%v] not found", name)
}

func (c *CommandFactory) Off(name string) string {
	if v, ok := c.commands[name]; ok {
		v.SetEnabled(false)
		if err := UpdateCmd(v); err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Command [%v] off", name)
	}
	return fmt.Sprintf("Command [%v] not found", name)
}

func (c *CommandFactory) SetInterval(name string, interval time.Duration) string {
	if v, ok := c.commands[name]; ok {
		v.SetIdleTime(interval)
		if err := UpdateCmd(v); err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Command [%v] interval: [%v]", name, interval)
	}
	return fmt.Sprintf("Command [%v] not found", name)
}

func (c *CommandFactory) GetJobs() ([]command.ICommand, error) {
	cursor, err := db.GetDb().Find("commands", bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		// cmd := command.Command{}
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}

		// Unmarshal the BSON document into the custom command object
		// Convert the result map to BSON representation
		resultBytes, err := bson.Marshal(result)
		if err != nil {
			return nil, err
		}

		cmd := &command.CustomCommand{}
		err = bson.Unmarshal(resultBytes, cmd)
		if err != nil {
			return nil, err
		}

		var iCmd interface{}
		switch cmd.Type {
		case command.ContractCall:
			iCmd = &contract.ContractCommand{}
		case command.Tomb:
			iCmd = &tomb.TombCommand{}
		case command.Balance:
			iCmd = &balance.BalanceCommand{}
		case command.BybitIdo:
			iCmd = &bybitido.IdoCommand{}
		default:
			log.GeneralLogger.Printf("unsupported cmd[%+v] ", cmd)
			continue
		}
		err = bson.Unmarshal(resultBytes, iCmd)
		if err != nil {
			return nil, err
		}

		_command, ok := iCmd.(command.ICommand)
		if !ok {
			return nil, fmt.Errorf("cannot typecast cmd[%v] to ICommand", cmd)
		}

		name := _command.GetName()
		c.commands[name] = _command
		b, err := json.MarshalIndent(_command, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("cannot MarshalIndent _command[%v]", _command)
		}
		log.GeneralLogger.Printf("Loaded Command [%v]\n", string(b))
	}

	return c.commands.ToActiveList(), nil
}
