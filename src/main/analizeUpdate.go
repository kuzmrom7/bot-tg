package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/bot-tg/src/store"
)

func (telegramBot *TelegramBot) analyzeUpdate(update tgbotapi.Update) {

	var (
		fromID int
		toID   int
		dateID int
	)

	var buttons = []tgbotapi.KeyboardButton{
		tgbotapi.KeyboardButton{Text: "Поехали " + TrainFrom},
		tgbotapi.KeyboardButton{Text: "/start "},
	}

	var button = []tgbotapi.KeyboardButton{
		tgbotapi.KeyboardButton{Text: "/start "},
	}

	for update := range telegramBot.Updates {
		var msg tgbotapi.MessageConfig
		log.Println("recived text: ", update.Message.Text)

		if (update.Message.MessageID-2 == fromID) {

			store.AddFrom(update.Message.Chat.ID, update.Message.Text)

			update.Message.Text = "Куда " + TraintTo
		}
		if (update.Message.MessageID-2 == toID) {
			store.AddTo(update.Message.Chat.ID, update.Message.Text)
			update.Message.Text = "Дата" + Date
		}
		if (update.Message.MessageID-2 == dateID) {
			store.AddDate(update.Message.Chat.ID, update.Message.Text)
			update.Message.Text = "Nice"
		}

		switch update.Message.Text {

		case "/start":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, нажми на кпопку,чтобы начать ? "+train)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

			store.CheckUser(update.Message.Chat.ID)

		case "Поехали " + TrainFrom:
			fromID = update.Message.MessageID
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Откуда едешь? "+TrainFrom)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(button)

		case "Куда " + TraintTo:
			toID = update.Message.MessageID
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Куда едешь? "+TraintTo)

		case "Дата" + Date:
			dateID = update.Message.MessageID
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Когда едешь? "+Date)

		case "Nice":
			dateID = update.Message.MessageID
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Все супер! "+Done)

		default:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие: ")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)
		}

		telegramBot.API.Send(msg)
	}
}
