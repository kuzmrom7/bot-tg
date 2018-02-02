package bot

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/bot-tg/src/store"
	"fmt"
)

func (telegramBot *TelegramBot) analyzeUpdate(update tgbotapi.Update) {

	var buttons = []tgbotapi.KeyboardButton{
		tgbotapi.KeyboardButton{Text: "Поехали " + TrainFrom},
		tgbotapi.KeyboardButton{Text: "СТАРТ "},
	}

	var button = []tgbotapi.KeyboardButton{
		tgbotapi.KeyboardButton{Text: "СТАРТ"},
	}

	for update := range telegramBot.Updates {
		var msg tgbotapi.MessageConfig

		log.Println("recived text: ", update.Message.Text)

		fromQuestions := store.FromQuestions(update.Message.Chat.ID)
		toQuestions := store.ToQuestions(update.Message.Chat.ID)
		dateQuestions := store.DateQuestions(update.Message.Chat.ID)


		if fromQuestions {
			store.WriteFromQuestions(update.Message.Chat.ID, false)
			store.AddFrom(update.Message.Chat.ID, update.Message.Text)
			update.Message.Text = "Куда " + TraintTo
		}

		if toQuestions {
			store.WriteToQuestions(update.Message.Chat.ID, false)
			store.AddTo(update.Message.Chat.ID, update.Message.Text)
			update.Message.Text = "Дата" + Date
		}

		if dateQuestions {
				store.WriteDateQuestions(update.Message.Chat.ID, false)
				store.AddDate(update.Message.Chat.ID, update.Message.Text)
				update.Message.Text = "Nice"
		}

		switch update.Message.Text {

		case "/start":
			messageStart := "Привет! Скорее всего ты мой друг и я скинул тебе бота чтобы потестить!" + WinkingFace + "\n" +
				"Пока что доступно только три города *Москва Санкт-Петербург* и *Орск*  ахах! \n Это все *ВПЕРЕД ТЕСТИТЬ* \n" +
				"`Нажми Поехали или СТАРТ` "
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, messageStart+train)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)
			msg.ParseMode = "markdown"

			store.CheckUser(update.Message.Chat.ID)

		case "СТАРТ":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, нажми на ПОЕХАЛИ,чтобы начать ? "+train)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

			store.CheckUser(update.Message.Chat.ID)

		case "Поехали " + TrainFrom:

			store.WriteFromQuestions(update.Message.Chat.ID, true)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Откуда едешь? "+WinkingFace)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(button)

		case "Куда " + TraintTo:

			store.WriteToQuestions(update.Message.Chat.ID, true)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Куда едешь? "+SmirkingFAce)

		case "Дата" + Date:

			store.WriteDateQuestions(update.Message.Chat.ID, true)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Когда едешь? (формат ввода `02.02.2018`)"+Date)

		case "Nice":
			msgAPI := AddData(update.Message.Chat.ID)

			msg = tgbotapi.NewMessage(update.Message.Chat.ID, msgAPI)
			msg.ParseMode = "markdown"
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

		default:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие: ")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)
		}

		telegramBot.API.Send(msg)
	}
}
