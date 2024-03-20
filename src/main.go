package main

import (
	"time"

	"github.com/hungdoo/bot/src/packages/db"
	"github.com/hungdoo/bot/src/packages/dotenv"
	"github.com/hungdoo/bot/src/packages/log"
	"github.com/hungdoo/bot/src/services/telecommands"
)

func main() {
	dotenv.InitEnv()

	_db := db.GetDb()
	defer _db.Close()

	commandService := telecommands.NewService()
	commandService.RegisterCommands()

	if dotenv.GetEnv("SERVICE_MODE") == "listener" {
		// go commandService.Work()
		go func() { // listener only refreshes jobs
			for {
				_, err := commandService.Factory.GetJobs()
				if err != nil {
					log.GeneralLogger.Printf("GetJobs err: [%s]", err)
				}
				time.Sleep(time.Second * 30)
			}
		}()
		commandService.ListenToCommand()
	} else {
		commandService.Work()
	}
}
