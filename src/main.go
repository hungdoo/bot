package main

import (
	"github.com/hungdoo/bot/src/packages/db"
	"github.com/hungdoo/bot/src/packages/dotenv"
	"github.com/hungdoo/bot/src/services/commands"
)

func main() {
	dotenv.InitEnv()

	_db := db.GetDb()
	defer _db.Close()

	commandService := commands.NewService()
	go commandService.Work()
	commandService.Run()
}
