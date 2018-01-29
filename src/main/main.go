package main

import "github.com/bot-tg/src/store"

func main() {

	var telegramBot TelegramBot

	store.New()
	telegramBot.Init()
	telegramBot.Start()

}
