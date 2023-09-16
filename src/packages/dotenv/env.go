package dotenv

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hungdoo/bot/src/packages/log"
	"github.com/joho/godotenv"
)

var BotEnvs map[string]string

func InitEnv() {
	path, _ := filepath.Abs(".")
	defaultEnv := os.Getenv("BOT_ENV")
	envPath := ""
	switch defaultEnv {
	case "dev":
		envPath = fmt.Sprintf("%s/.env.%s", path, defaultEnv)
	case "prod":
		envPath = "/data/.env.prod"
	default:
		envPath = fmt.Sprintf("%s/.env", path)
	}

	var err error
	BotEnvs, err = godotenv.Read(envPath)
	if err != nil {
		log.GeneralLogger.Fatalf("Error loading %s file: %s", envPath, err)
	}

	for k, v := range BotEnvs {
		log.GeneralLogger.Printf("%s:%s\n", k, v)
	}
}

func Abs(s string) {
	panic("unimplemented")
}

func GetEnv(k string) (v string) {
	return BotEnvs[k]
}
