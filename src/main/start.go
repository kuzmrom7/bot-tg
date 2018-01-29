package main

import "gopkg.in/telegram-bot-api.v4"

func (telegramBot *TelegramBot) Start() {
	for update := range telegramBot.Updates {
		if update.Message != nil {
			telegramBot.analyzeUpdate(update)
		}
		if update.CallbackQuery != nil {
			class := update.CallbackQuery.Data
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, class+" Okey")
			telegramBot.API.Send(msg)
		}
	}
}
