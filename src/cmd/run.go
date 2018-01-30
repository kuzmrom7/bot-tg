package main

import (
	"github.com/bot-tg/src/store"
	"github.com/bot-tg/src/bot"
)

func main() {

	var telegramBot bot.TelegramBot

	store.New()
	telegramBot.Init()
	telegramBot.Start()

}
