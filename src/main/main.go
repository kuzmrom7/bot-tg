package main

import (
	"os"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
)

var buttons = []tgbotapi.KeyboardButton{
	tgbotapi.KeyboardButton{Text: "Москва"},
	tgbotapi.KeyboardButton{Text: "Санкт-Петербург"},
}

var (
	train = "\xF0\x9F\x9A\x83"
	help = "\xF0\x9F\x99\x8F"
)

const WebHookURL = "https://bot-kuzmen.herokuapp.com/"


func main() {
	port := os.Getenv("PORT")
	bot, err := tgbotapi.NewBotAPI("475819101:AAFuuJ51XbSkj3vd91U0aUHh2Gnk_CpwUhA")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Autorizen on account %s", bot.Self.UserName)

	//Install WebHook

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebHookURL))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+port, nil)

	for update := range updates {
		var msg tgbotapi.MessageConfig
		log.Println("recived text: ", update.Message.Text)

		switch update.Message.Text {
		case "/start":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, откуда поедешь? "+train)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)
		default:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Loool "+update.Message.Text+help)
		}


		bot.Send(msg)
	}
}
