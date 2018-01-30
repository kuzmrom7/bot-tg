package bot

func (telegramBot *TelegramBot) Start() {
	for update := range telegramBot.Updates {
		if update.Message != nil {
			telegramBot.analyzeUpdate(update)
		}
	}
}

