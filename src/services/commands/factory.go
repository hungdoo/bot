package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/hungdoo/bot/src/packages/contract"
	"github.com/hungdoo/bot/src/types"
)

type CommandFactory struct {
	commands map[string]types.ICommand
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
		return fmt.Sprintf("Command [%v] updated", name)
	} else {
		switch name {
		case "callContract":
			newCommand := &contract.ContractCommand{
				Command: types.Command{
					Name:     name,
					Data:     data,
					Enabled:  true,
					IdleTime: time.Second * 30,
				},
			}
			if err := newCommand.SetData(data); err != nil {
				return err.Error()
			}
			c.commands[name] = newCommand
		default:
			return fmt.Sprintf("Invalid Command [%v] added", name)
		}
		return fmt.Sprintf("Command [%v] added", name)
	}
}

func (c *CommandFactory) Remove(name string) string {
	if _, ok := c.commands[name]; ok {
		delete(c.commands, name)
		return fmt.Sprintf("Command [%v] removed", name)
	}
	return fmt.Sprintf("Command [%v] not found", name)
}

func (c *CommandFactory) Exec(name string) (string, error) {
	if v, ok := c.commands[name]; ok {
		var cmd types.ICommand = v
		result, err := cmd.Execute(true)
		return result, err
	}
	return "", fmt.Errorf("command [%v] not found", name)
}

func (c *CommandFactory) List() string {
	cmdTxt := []string{}
	for _, cmd := range c.commands {
		if cmd.IsEnabled() {
			cmdTxt = append(cmdTxt, fmt.Sprintf("<strong>%v: data(%v)</strong>", cmd.GetName(), strings.Join(cmd.GetData(), ", ")))
		} else {
			cmdTxt = append(cmdTxt, fmt.Sprintf("%v: data(%v)", cmd.GetName(), strings.Join(cmd.GetData(), ", ")))
		}
	}
	return strings.Join(cmdTxt, "\n")
}

func (c *CommandFactory) On(name string) string {
	if v, ok := c.commands[name]; ok {
		v.SetEnabled(true)
		return fmt.Sprintf("Command [%v] on", name)
	}
	return fmt.Sprintf("Command [%v] not found", name)
}

func (c *CommandFactory) Off(name string) string {
	if v, ok := c.commands[name]; ok {
		v.SetEnabled(false)
		return fmt.Sprintf("Command [%v] off", name)
	}
	return fmt.Sprintf("Command [%v] not found", name)
}

func (c *CommandFactory) GetJobs() []types.ICommand {
	actives := make([]types.ICommand, 0, len(c.commands))

	for _, v := range c.commands {
		if v.IsEnabled() {
			actives = append(actives, v)
		}
	}
	return actives
}
