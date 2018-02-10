package main

import (
	"github.com/kuzmrom7/bot-tg/src/store"
	"github.com/kuzmrom7/bot-tg/src/bot"
)

func main() {

	var telegramBot bot.TelegramBot

	store.New()
	telegramBot.Init()
	telegramBot.Start()

}
