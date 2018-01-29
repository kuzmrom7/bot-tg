package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"os"
	"log"
	"net/http"
)

const WebHookURL = "https://bot-kuzmen.herokuapp.com/"

type TelegramBot struct {
	API                   *tgbotapi.BotAPI        // API телеграмма
	Updates               tgbotapi.UpdatesChannel // Канал обновлений
	ActiveContactRequests []int64                 // ID чатов, от которых мы ожидаем номер
}



func (telegramBot *TelegramBot)Init(){
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
	telegramBot.Updates = bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+port, nil)
}
