package telegram

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hungdoo/bot/src/packages/dotenv"
)

var _bot *tgbotapi.BotAPI
var _updates tgbotapi.UpdatesChannel
var WHITELISTED = make(map[string]bool)

func loadWhitelist() {
	whitelists := strings.Split(dotenv.GetEnv("WHITELISTS"), ",")
	for _, wl := range whitelists {
		if len(wl) != 0 {
			WHITELISTED[wl] = true
		}
	}
}
func IsWhitelisted(userName string) bool {
	_, allowed := WHITELISTED[userName]
	return allowed
}

func ReportInvalidAccess(fromUser string) {
	reportChatId := dotenv.GetEnv("REPORT_CHAT_ID")
	chatId, err := strconv.ParseInt(reportChatId, 10, 64)
	if err != nil {
		return
	}
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("[!!]Access from invalid user [%s]", fromUser))

	GetBot().Send(msg)
	time.Sleep(time.Millisecond * 500)
}

func GetBot() *tgbotapi.BotAPI {
	if _bot == nil {
		bot, err := tgbotapi.NewBotAPI(dotenv.BotEnvs["BOT_TELEGRAM_TOKEN"])
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			return nil
		}
		loadWhitelist()
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
