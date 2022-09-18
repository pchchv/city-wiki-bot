package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func bot() {
	bot, err := tgbotapi.NewBotAPI(getEnvValue("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	// Set the update time
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	// Getting updates from the bot
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
	}
}
