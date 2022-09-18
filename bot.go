package main

import (
	"fmt"
	"log"
	"reflect"

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
		// Check that a text message was sent by the user
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				text := fmt.Sprintf(
					"Hi %v!\nI'm a City Wiki Bot.\nSend me the name of the country and I'll reply with an emoji with its flag and a link to the Wikipedia page.",
					update.Message.From,
				)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)
			}
		}
	}
}
