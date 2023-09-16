package main

import (
	"github.com/hungdoo/bot/src/packages/db"
	u "github.com/hungdoo/bot/src/packages/utils"
	"github.com/hungdoo/bot/src/services/commands"
)

func main() {
	u.InitEnv()

	_db := db.GetDb()
	defer _db.Close()

	commandService := commands.NewService()
	go commandService.Work()
	commandService.Run()
}
