package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
)


func (telegramBot *TelegramBot) analyzeUpdate(update tgbotapi.Update) {

	classes := []string{
		"Москва",
		"СПб",
	}

	for update := range telegramBot.Updates {
		var msg tgbotapi.MessageConfig
		log.Println("recived text: ", update.Message.Text)


			switch update.Message.Text {

			case "/start":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, откуда поедешь? "+train)

				keyboard := tgbotapi.InlineKeyboardMarkup{}
				for _, class := range classes {
					var row []tgbotapi.InlineKeyboardButton
					btn := tgbotapi.NewInlineKeyboardButtonData(class, class)
					row = append(row, btn)
					keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
				}

				msg.ReplyMarkup = keyboard

			default:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			}


		telegramBot.API.Send(msg)
	}
}
