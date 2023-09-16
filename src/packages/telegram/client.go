package telegram

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hungdoo/bot/src/packages/dotenv"
)

var _bot *tgbotapi.BotAPI
var _updates tgbotapi.UpdatesChannel

func GetBot() *tgbotapi.BotAPI {
	if _bot == nil {
		bot, err := tgbotapi.NewBotAPI(dotenv.BotEnvs["BOT_TELEGRAM_TOKEN"])
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			return nil
		}
		_bot = bot
	}
	return _bot
}

func GetUpdates() (tgbotapi.UpdatesChannel, error) {
	bot := GetBot()

	if _updates == nil {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)
		if err != nil {
			return nil, fmt.Errorf("%v", err.Error())
		}

		time.Sleep(time.Millisecond * 500)
		updates.Clear()
		_updates = updates
	}

	return _updates, nil
}
