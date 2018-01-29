package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/bot-tg/src/store"
)


func (telegramBot *TelegramBot) analyzeUpdate(update tgbotapi.Update) {


	for update := range telegramBot.Updates {
		var msg tgbotapi.MessageConfig
		log.Println("recived text: ", update.Message.Text)
		log.Println("id --> ", update.Message.Chat.ID)

			switch update.Message.Text {

			case "/start":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, откуда поедешь? "+train)

				store.CreateNote(update.Message.Chat.ID)

			default:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			}


		telegramBot.API.Send(msg)
	}
}
