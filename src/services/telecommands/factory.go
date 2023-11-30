package telecommands

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hungdoo/bot/src/packages/command"
	"github.com/hungdoo/bot/src/packages/contract"
	"github.com/hungdoo/bot/src/packages/db"
	"github.com/hungdoo/bot/src/packages/debank"
	"github.com/hungdoo/bot/src/packages/interfaces"
	"github.com/hungdoo/bot/src/packages/log"
	"go.mongodb.org/mongo-driver/bson"
)

type CommandMap map[string]interfaces.ICommand

func (cm *CommandMap) Search(nameSubString string) []interfaces.ICommand {
	list := make([]interfaces.ICommand, 0, len(*cm))
	for _, v := range *cm {
		if strings.Contains(v.GetName(), nameSubString) {
			list = append(list, v)
		}
	}
	return list
}
func (cm *CommandMap) ToList() []interfaces.ICommand {
	list := make([]interfaces.ICommand, 0, len(*cm))
	for _, v := range *cm {
		list = append(list, v)
	}
	return list
}
func (cm *CommandMap) ToActiveList() []interfaces.ICommand {
	actives := make([]interfaces.ICommand, 0, len(*cm))
	for _, v := range *cm {
		if v.IsEnabled() {
			actives = append(actives, v)
		}
	}
	return actives
}

type CommandFactory struct {
	commands          CommandMap
	lastRefreshedAt   time.Time
	refreshDbInterval time.Duration
}

func NewCommadFactory() CommandFactory {
	return CommandFactory{commands: map[string]interfaces.ICommand{}, lastRefreshedAt: time.Now(), refreshDbInterval: 3 * time.Minute}
}

func (c *CommandFactory) Add(messages []string) string {
	if len(messages) < 2 {
		return "Add needs at least 2 params"
	}
	name := messages[0]
	data := messages[1:]
	if v, ok := c.commands[name]; ok {
		if err := v.SetData(data); err != nil {
			return err.Error()
		}
		filter := bson.M{"name": v.GetName()}
		update := bson.M{"$set": bson.M{"data": v.GetData()}}
		if err := db.GetDb().Update("commands", filter, update); err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Command [%v] updated", name)
	} else {
		var newCommand interfaces.ICommand

		switch command.IsType(name) {
		case command.ContractCall:
			newCommand = &contract.Command{
				Command: command.Command{
					Name:     name,
					Enabled:  true,
					IdleTime: time.Second * 60,
				},
			}
		case command.Debank:
			newCommand = &debank.Command{
				Command: command.Command{
					Name:     name,
					Enabled:  true,
					IdleTime: time.Second * 60,
				},
			}
		}
		if newCommand == nil {
			return fmt.Sprintf("Command [%v:%v] failed to add", name, command.IsType(name))
		}

		if err := newCommand.SetData(data); err != nil {
			return err.Error()
		}
		if err := newCommand.SetType(name); err != nil {
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

func (c *CommandFactory) Exec(name string) (res string, err error) {
	searchedList := c.commands.Search(name)
	if len(searchedList) == 0 {
		return "", fmt.Errorf("command [%v] not found", name)
	}

	var executed []string
	for _, cmd := range searchedList {
		result, err := cmd.Execute(true)
		if err != nil {
			log.GeneralLogger.Printf("Job [%s] exec failed: [%s]", cmd.GetName(), err)
			continue
		}
		executed = append(executed, result)
	}
	return string(strings.Join(executed, "\n")), nil
}

func (c *CommandFactory) List() string {
	cmdTxt := []string{}
	for _, cmd := range c.commands {
		if cmd.IsEnabled() {
			cmdTxt = append(cmdTxt, fmt.Sprintf("[+] %v", cmd.GetName()))
		} else {
			cmdTxt = append(cmdTxt, fmt.Sprintf("[-] %v", cmd.GetName()))
		}
	}
	return strings.Join(cmdTxt, "\n")
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

func (c *CommandFactory) GetJobs() ([]interfaces.ICommand, error) {
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
			var _command interfaces.ICommand
			switch cmd.Type {
			case command.ContractCall:
				_command = &contract.Command{
					Command: cmd,
				}
			case command.Debank:
				_command = &debank.Command{
					Command: cmd,
				}
			}
			name := cmd.GetName()
			c.commands[name] = _command
			log.GeneralLogger.Printf("Loaded Command [%v]\n", name)
		}
	}

	return c.commands.ToActiveList(), nil
}
