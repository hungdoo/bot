package main

import (
	"github.com/hungdoo/bot/src/packages/db"
	"github.com/hungdoo/bot/src/packages/dotenv"
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
		commandService.ListenToCommand()
	} else {
		commandService.Work()
	}
}
