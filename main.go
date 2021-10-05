package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var BotToken = "45charlongdata"

func main() {
	//for range time.Tick(time.Minute * 10) {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	var telegramID int64 = 012345
	mesaj := "Bu bir mesaj örneğidir."
	msg := tgbotapi.NewMessage(telegramID, mesaj)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.DisableWebPagePreview = true
	bot.Send(msg)
	//}
}
