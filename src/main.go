package main

import (
	u "github.com/hungdoo/bot/src/packages/utils"
	"github.com/hungdoo/bot/src/services/commands"
	"github.com/hungdoo/bot/src/types"
)

func main() {
	u.InitEnv()
	u.GeneralLogger.Println("BOT Envs")
	for k, v := range u.BotEnvs {
		u.GeneralLogger.Printf("%s:%s\n", k, v)
	}

	var commandService types.IService = commands.NewService()
	go commandService.Work()
	commandService.Run()
}
