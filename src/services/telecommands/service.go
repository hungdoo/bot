package telecommands

import (
	"context"
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hungdoo/bot/src/packages/cmdparser"
	"github.com/hungdoo/bot/src/packages/db"
	"github.com/hungdoo/bot/src/packages/interfaces"
	"github.com/hungdoo/bot/src/packages/log"
	"github.com/hungdoo/bot/src/packages/telegram"
	"github.com/urfave/cli"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommandService struct {
	interfaces.IService
	Factory CommandFactory
	ChatID  int64
	Parser  *cli.App
}

func (s *CommandService) RegisterCommands() {
	s.Parser.Commands = []cli.Command{
		{
			Name:      "add",
			Aliases:   []string{"a"},
			Usage:     "add a task to the list",
			UsageText: "callContract-xyz, rpc, contract address, method (axy(address)(uint256), params(pr1;pr2), value index, margin(1%), usdcPrice;precision",
			Action: func(c *cli.Context) error {
				fmt.Fprintln(s.Parser.Writer, s.Factory.Add(c.Args()))
				return nil
			},
			// TODO: split Action into SubCommands: callContract, debank
			// Usage contract: rpc, contractAddr, method, params<;>, valueIdx, marginStr, precisionStr
			// Usage debank: ...
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Flags: []cli.Flag{
				cli.StringFlag{Name: "task,t", Required: true},
			},
			Usage: "rm a task",
			Action: func(c *cli.Context) error {
				task := c.String("task")
				fmt.Fprintln(s.Parser.Writer, s.Factory.Remove(task))
				return nil
			},
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Flags: []cli.Flag{
				cli.StringFlag{Name: "task,t", Required: true},
			},
			Usage: "show a task details",
			Action: func(c *cli.Context) error {
				task := c.String("task")
				fmt.Fprintln(s.Parser.Writer, s.Factory.Show(task))
				return nil
			},
		},
		{
			Name:    "execute",
			Aliases: []string{"exec"},
			Usage:   "exec a task",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "task,t", Required: true},
			},
			Action: func(c *cli.Context) error {
				task := c.String("task")
				res, err := s.Factory.Exec(task)
				if err != nil {
					fmt.Fprintln(s.Parser.Writer, err.Error())
					return err
				}
				fmt.Fprintln(s.Parser.Writer, res)
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "list tasks",
			Action: func(c *cli.Context) error {
				fmt.Fprintln(s.Parser.Writer, s.Factory.List())
				return nil
			},
		},
		{
			Name:    "on",
			Aliases: []string{"on"},
			Usage:   "enable a task",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "task,t", Required: true},
			},
			Action: func(c *cli.Context) error {
				task := c.String("task")
				fmt.Fprintln(s.Parser.Writer, s.Factory.On(task))
				return nil
			},
		},
		{
			Name:    "off",
			Aliases: []string{"off"},
			Usage:   "disable a task",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "task,t", Required: true},
			},
			Action: func(c *cli.Context) error {
				task := c.String("task")
				fmt.Fprintln(s.Parser.Writer, s.Factory.Off(task))
				return nil
			},
		},
		{
			Name:    "setInterval",
			Aliases: []string{"si"},
			Usage:   "set new run interval for task in s",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "task,t", Required: true},
				cli.Int64Flag{Name: "duration, d", Required: true},
			},
			Action: func(c *cli.Context) error {
				task := c.String("task")
				duration := time.Duration(c.Int64("duration")) * time.Second
				fmt.Fprintln(s.Parser.Writer, s.Factory.SetInterval(task, duration))
				return nil
			},
		},
	}
}

func (c *CommandService) process(message string) string {
	ret := "No action"
	messages := strings.Split(strings.TrimSpace(message), " ")
	if err := c.Parser.Run(append([]string{"tele"}, messages...)); err != nil {
		ret = err.Error()
	} else {
		ret = cmdparser.GetOutput()
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
				// exec seccessfully -> update prev in db
				filter := bson.M{"name": j.GetName()}
				update := bson.M{"$set": bson.M{"prev": j.GetPrev().String()}} // mongo not support decimal type
				if err := db.GetDb().Update("commands", filter, update); err != nil {
					log.GeneralLogger.Printf("Job [%s] update db failed: [%s]", j.GetName(), err)
					continue
				}
				if result != "" {
					msg := tgbotapi.NewMessage(c.ChatID, result)
					msg.ParseMode = ""
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
		msg.ParseMode = ""

		telegram.GetBot().Send(msg)
		if c.ChatID == 0 {
			c.ChatID = update.Message.Chat.ID
		}
	}
	return nil
}

func NewService() *CommandService {
	parser := cmdparser.GetParser()
	keys := bson.M{
		"name": 1, // index in ascending order; -1 for descending order
	}
	index := mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetUnique(true),
	}
	_, err := db.GetDb().GetCollection("commands").Indexes().CreateOne(context.Background(), index)
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}

	return &CommandService{Factory: CommandFactory{commands: map[string]interfaces.ICommand{}, lastRefreshedAt: time.Now(), refreshDbInterval: 3 * time.Minute}, Parser: parser}
}
