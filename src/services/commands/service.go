package commands

import (
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hungdoo/bot/src/packages/interfaces"
	"github.com/hungdoo/bot/src/packages/log"
	"github.com/hungdoo/bot/src/packages/telegram"
)

type CommandService struct {
	interfaces.IService
	Factory CommandFactory
	ChatID  int64
}

func (c *CommandService) process(message string) string {
	ret := "No action"
	messages := strings.Split(strings.TrimSpace(message), " ")
	if len(messages) > 0 {
		switch messages[0] {
		case "add":
			ret = c.Factory.Add(messages[1:])
		case "rm":
			if len(messages) < 2 {
				ret = "Rm command needs `name` arg"
				break
			}
			ret = c.Factory.Remove(messages[1])
		case "exec":
			if len(messages) < 2 {
				ret = "Exec command needs `name` arg"
				break
			}
			result, err := c.Factory.Exec(messages[1])
			if err != nil {
				ret = err.Error()
				break
			}
			if result != "" {
				ret = ""
				msg := tgbotapi.NewMessage(c.ChatID, result)
				msg.ParseMode = "HTML"
				telegram.GetBot().Send(msg)
			}
		case "ls":
			ret = c.Factory.List()
		case "on":
			if len(messages) < 2 {
				ret = "On command needs `name` arg"
				break
			}
			ret = c.Factory.On(messages[1])
		case "off":
			if len(messages) < 2 {
				ret = "Off command needs `name` arg"
				break
			}
			ret = c.Factory.Off(messages[1])
		}
	} else {
		ret = "Empty message"
	}
	log.GeneralLogger.Printf("Process message[%s] & ret[%s]\n", message, ret)
	return ret
}

// Work all commands in intervals
func (c *CommandService) Work() {
	for {
		jobs, err := c.Factory.GetJobs()
		if err != nil {
			log.GeneralLogger.Printf("GetJobs err: [%s]", err)
		}
		for _, j := range jobs {
			if !j.IsIdle() {
				result, err := j.Execute(false)
				j.SetExecutedTime(time.Now())
				if err != nil {
					log.GeneralLogger.Printf("Job [%s] exec failed: [%s]", j.GetName(), err)
					continue
				}
				if result != "" {
					msg := tgbotapi.NewMessage(c.ChatID, result)
					msg.ParseMode = "HTML"
					telegram.GetBot().Send(msg)
				}
			}
		}
		time.Sleep(time.Millisecond * 5000)
	}
}

func (c *CommandService) Run() error {
	telegramUpdateChan, err := telegram.GetUpdates()
	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}
	for update := range telegramUpdateChan {
		if update.Message == nil {
			continue
		}

		log.GeneralLogger.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, c.process(update.Message.Text))
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "HTML"

		telegram.GetBot().Send(msg)
		if c.ChatID == 0 {
			c.ChatID = update.Message.Chat.ID
		}
	}
	return nil
}

func NewService() *CommandService {
	return &CommandService{Factory: CommandFactory{commands: map[string]interfaces.ICommand{}}}
}
