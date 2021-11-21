package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

var BotEnvs = make(map[string]string)

func InitEnv() {
	path, _ := filepath.Abs(".")
	defaultEnv := os.Getenv("BOT_ENV")
	switch defaultEnv {
	case "dev", "prod":
		envPath := fmt.Sprintf("%s/.env.%s", path, defaultEnv)
		err := godotenv.Load(envPath)
		if err != nil {
			log.Fatalf("Error loading %s file", envPath)
		}
	default:
		envPath := fmt.Sprintf("%s/.env", path)
		err := godotenv.Load(envPath)
		if err != nil {
			log.Fatalf("Error loading %s file", envPath)
		}
	}

	// Double check ENV with BOT_ prefixes
	for _, env := range os.Environ() {
		envPair := strings.SplitN(env, "=", 2)
		key := envPair[0]
		value := envPair[1]

		if strings.Contains(key, "BOT_") {
			BotEnvs[key] = value
		}
	}
}

func Abs(s string) {
	panic("unimplemented")
}
