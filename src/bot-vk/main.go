package main

import (
	"log"
	"github.com/Dimonchik0036/vk-api"
	"github.com/bot-tg/src/bot"
)

func main() {
	client, err := vkapi.NewClientFromToken("09a787cfe4e3b13f2c620ce0f059b5409fcc422dfe7800cdb1fa977b7867738e831882e858473b0d60620")
	if err != nil {
		log.Panic(err)
	}

	client.Log(true)

	if err := client.InitLongPoll(0, 2); err != nil {
		log.Panic(err)
	}

	updates, _, err := client.GetLPUpdatesChan(100, vkapi.LPConfig{25, vkapi.LPModeAttachments})
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil || !update.IsNewMessage() || update.Message.Outbox(){
			continue
		}

		log.Printf("%s", update.Message.String())

		if update.Message.Text == "/start" {
			client.SendMessage(vkapi.NewMessage(vkapi.NewDstFromUserID(update.Message.FromID), "Hello!" + bot.WinkingFace))
		}


	}
}