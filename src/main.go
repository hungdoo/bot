package main

import (
	u "github.com/hungdoo/bot/src/packages/utils"
)

func main() {
	u.InitEnv()
	u.GeneralLogger.Println("BOT Envs")
	for k, v := range u.BotEnvs {
		u.GeneralLogger.Printf("%s:%s\n", k, v)
	}
	return
}
