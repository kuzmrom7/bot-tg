package main

import "gopkg.in/telegram-bot-api.v4"

func HandlerMessage(message string, id int64) (tgbotapi.MessageConfig, error){
	var msg tgbotapi.MessageConfig

	switch message {
	case "Hello":
		msg = tgbotapi.NewMessage(id, "Hi, how can i help you?")
	default:
		msg = tgbotapi.NewMessage(id, "Loool"+message)
	}

	return msg, nil
}
