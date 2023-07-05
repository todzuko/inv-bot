package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/todzuko/inv-bot/api/routes"
	"log"
	"os"
)

func Connect() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60

	updates := bot.GetUpdatesChan(upd)
	for update := range updates {
		if update.Message != nil {
			msgText := update.Message.Text
			res := routes.GetResponse(msgText, update.Message.Chat.ID)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, res)
			bot.Send(msg)
		}
	}
}
