package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/todzuko/inv-bot/api/routes"
	"log"
	"os"
)

func Connect() {
	fmt.Println(os.Getenv("BOT_TOKEN"))
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
			res := `Not a command, please try again`
			msgText := update.Message.Text
			if string(msgText[0]) == "/" {
				res = routes.GetResponse(msgText)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, res)
			bot.Send(msg)
		}
	}
}
