package telecommands

import (
	"context"
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hungdoo/bot/src/common"
	"github.com/hungdoo/bot/src/packages/cmdparser"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/db"
	"github.com/hungdoo/bot/src/packages/interfaces"
	"github.com/hungdoo/bot/src/packages/log"
	"github.com/hungdoo/bot/src/packages/telegram"
	"github.com/hungdoo/bot/src/packages/utils"
	lock "github.com/square/mongo-lock"
	"github.com/urfave/cli"
)

// @dev external facing controller
// Responsibility: listen to command from tele, executes & reports
type CommandService struct {
	interfaces.IService
	Factory CommandFactory
	Parser  *cli.App
	LockCli *lock.Client
}

func (s *CommandService) RegisterCommands() {
	s.Parser.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Subcommands: []cli.Command{
				{
					Name:      "call",
					UsageText: "name, rpc, contract address, method (axy(address)(uint256), params(pr1;pr2), value index, margin(1%), usdcPrice;precision",
					Action: func(c *cli.Context) error {
						fmt.Fprintln(s.Parser.Writer, s.Factory.Add(command.ContractCall, c.Args()))
						return nil
					},
				},
				{
					Name:      "tomb",
					UsageText: "name, rpc, contract, up, pkIdx, k",
					Action: func(c *cli.Context) error {
						fmt.Fprintln(s.Parser.Writer, s.Factory.Add(command.Tomb, c.Args()))
						return nil
					},
				},
				{
					Name:      "balance",
					UsageText: "name, rpc, ...wallets",
					Action: func(c *cli.Context) error {
						fmt.Fprintln(s.Parser.Writer, s.Factory.Add(command.Balance, c.Args()))
						return nil
					},
				},
				{
					Name:      "bybit",
					UsageText: "name",
					Action: func(c *cli.Context) error {
						fmt.Fprintln(s.Parser.Writer, s.Factory.Add(command.BybitIdo, c.Args()))
						return nil
					},
				},
			},
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
			Subcommands: []cli.Command{
				{
					Name: "call",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "task,t", Required: true},
					},
					Action: func(c *cli.Context) error {
						task := c.String("task")
						res, err := s.Factory.Exec(command.ContractCall, task, "")
						if err != nil {
							fmt.Fprintln(s.Parser.Writer, err.Error())
							return err
						}
						fmt.Fprintln(s.Parser.Writer, res)
						return nil
					},
				},
				{
					Name: "balance",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "task,t", Required: true},
					},
					Action: func(c *cli.Context) error {
						task := c.String("task")
						res, err := s.Factory.Exec(command.Balance, task, "")
						if err != nil {
							fmt.Fprintln(s.Parser.Writer, err.Error())
							return err
						}
						fmt.Fprintln(s.Parser.Writer, res)
						return nil
					},
				},
				{
					Name: "tomb",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "task,t", Required: true},
						cli.StringFlag{Name: "subcommand,sc", Usage: "stats, clear, claim, cronjob, default", Value: "stats"},
					},
					Action: func(c *cli.Context) error {
						task := c.String("task")
						subcommand := c.String("subcommand")
						res, err := s.Factory.Exec(command.Tomb, task, subcommand)
						if err != nil {
							fmt.Fprintln(s.Parser.Writer, err.Error())
							return err
						}
						fmt.Fprintln(s.Parser.Writer, res)
						return nil
					},
				},
				{
					Name: "bybit",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "task,t", Required: true},
						cli.StringFlag{Name: "subcommand,sc", Usage: "latest, all, default", Value: "latest"},
					},
					Action: func(c *cli.Context) error {
						task := c.String("task")
						subcommand := c.String("subcommand")
						res, err := s.Factory.Exec(command.BybitIdo, task, subcommand)
						if err != nil {
							fmt.Fprintln(s.Parser.Writer, err.Error())
							return err
						}
						fmt.Fprintln(s.Parser.Writer, res)
						return nil
					},
				},
			},
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all,a", Usage: "show all"},
			},
			Usage: "list tasks",
			Action: func(c *cli.Context) error {
				showAll := c.Bool("all")
				fmt.Fprintln(s.Parser.Writer, s.Factory.List(showAll))
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
				cli.StringFlag{Name: "duration, d", Required: true, Usage: "30s, 1m, 2h"},
			},
			Action: func(c *cli.Context) error {
				task := c.String("task")
				duration, err := utils.ParseCustomDuration(c.String("duration"))
				if err != nil {
					fmt.Fprintln(s.Parser.Writer, err.Error())
					return err
				}
				fmt.Fprintln(s.Parser.Writer, s.Factory.SetInterval(task, duration))
				return nil
			},
		},
	}
}

func (c *CommandService) process(message string) string {
	// Refresh jobs
	_, err := c.Factory.GetJobs()
	if err != nil {
		return fmt.Sprintf("GetJobs err: [%s]", err)
	}

	messages := strings.Split(strings.TrimSpace(message), " ")
	if err = c.Parser.Run(append([]string{"tele"}, messages...)); err != nil {
		return fmt.Sprintf("Parser run err: %s", err)
	}
	return cmdparser.GetOutput()
}

// Work all commands in intervals
func (c *CommandService) Work() {
	for {
		jobs, err := c.Factory.GetJobs()
		if err != nil {
			log.GeneralLogger.Printf("GetJobs err: [%s]", err)
		}
		var results []string
		for _, j := range jobs {
			// Create an exclusive lock on a resource.
			lockId := "mutexLock" + j.GetName()
			err = c.LockCli.XLock(context.Background(), j.GetName(), lockId, lock.LockDetails{})
			if err != nil {
				log.GeneralLogger.Printf("Job [%s] XLock err: [%s]", j.GetName(), err.Error())
				j.SetError(err.Error())
				continue
			}

			if !j.IsIdle() {
				result, execErr := j.Execute(false, "")
				log.GeneralLogger.Printf("[%s] execution: [%s]", j.GetName(), result)
				j.SetExecutedTime(time.Now())
				if execErr != nil && execErr.Level >= common.Error {
					log.GeneralLogger.Printf("Job [%s] exec failed: [%s]", j.GetName(), execErr.Error())
					if execErr.Level >= common.Critical { // should report asap
						off := c.Factory.Off(j.GetName())
						results = append(results, fmt.Sprintf("%v failed with CRIT[%s]: off[%s]", j.GetName(), execErr.Error(), off))
					} else {
						j.SetError(execErr.Error()) // log for ls cmd
						// results = append(results, fmt.Sprintf("%v failed with [%s:%s]", j.GetName(), execErr.Level, execErr.Error()))
					}
				}

				// record result & info error for logging with tele.List cmd, no realtime report
				if result != "" {
					j.SetDisplayMsg(result)
					results = append(results, fmt.Sprintf("[%s]\n%s", j.GetName(), result))
				}
				// exec seccessfully -> update db
				UpdateCmd(j)
			}

			// Unlock the resource
			_, err = c.LockCli.Unlock(context.Background(), lockId)
			if err != nil {
				log.GeneralLogger.Printf("XLock Unlock err: [%s]", err)
				results = append(results, fmt.Sprintf("%v XLock Unlock failed with [%s]", j.GetName(), err))
			}
			log.GeneralLogger.Printf("[%s] XLock Release successfully", j.GetName())
		}

		if len(results) != 0 {
			reportChatId, err := telegram.GetReportChatId()
			if err != nil {
				log.GeneralLogger.Printf("[work] get report id failed: [%s]", err)
			}
			msg := tgbotapi.NewMessage(reportChatId, strings.Join(results, "\n"))
			msg.ParseMode = ""
			telegram.GetBot().Send(msg)
		}
		time.Sleep(time.Second * 30)
	}
}

func (c *CommandService) ListenToCommand() error {
	telegramUpdateChan, err := telegram.GetUpdates()
	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}
	for update := range telegramUpdateChan {
		if update.Message == nil {
			continue
		}

		fromUser := update.Message.From.UserName

		if !telegram.IsWhitelisted(fromUser) {
			telegram.ReportInvalidAccess(fromUser)
			continue
		}
		reportChatId, err := telegram.GetReportChatId()
		if err != nil {
			log.ErrorLogger.Print(err)
		}
		log.GeneralLogger.Printf("[%s:%d] %s", fromUser, reportChatId, update.Message.Text)
		msg := tgbotapi.NewMessage(reportChatId, c.process(update.Message.Text))
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = ""

		telegram.GetBot().Send(msg)
	}
	return nil
}

func NewService() *CommandService {
	parser := cmdparser.GetParser()
	// keys := bson.M{
	// 	"_id": 1, // index in ascending order; -1 for descending order
	// }
	// index := mongo.IndexModel{
	// 	Keys:    keys,
	// 	Options: options.Index().SetUnique(true),
	// }
	// _, err := db.GetDb().GetCollection("commands").Indexes().CreateOne(context.Background(), index)
	// if err != nil {
	// 	log.ErrorLogger.Fatal(err)
	// }

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	collection := db.GetDb().GetCollection("commands_lock")

	// Create a MongoDB lock client.
	cli := lock.NewClient(collection)

	// Create the required and recommended indexes.
	cli.CreateIndexes(ctx)

	return &CommandService{Factory: NewCommandFactory(), Parser: parser, LockCli: cli}
}
