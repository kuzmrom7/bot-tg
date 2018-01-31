package bot

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
		tgbotapi.KeyboardButton{Text: "СТАРТ "},
	}

	var button = []tgbotapi.KeyboardButton{
		tgbotapi.KeyboardButton{Text: "СТАРТ"},
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
			messageStart := "Привет! Скорее всего ты мой друг и я скинул тебе бота чтобы потестить!"+WinkingFace+"\n" +
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
			fromID = update.Message.MessageID
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Откуда едешь? "+WinkingFace)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(button)

		case "Куда " + TraintTo:
			toID = update.Message.MessageID
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Куда едешь? "+SmirkingFAce)

		case "Дата" + Date:
			dateID = update.Message.MessageID
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
