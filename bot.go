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
	log.Printf("Authorized on account %s\n", bot.Self.UserName)
	log.Println("Bot is running")
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
					"Привет %v!\nЯ City Wiki Bot.\nОтправь мне название страны, а я отвечу эмодзи с ее флагом и ссылкой на страницу страны в Википедии",
					update.Message.From,
				)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)
			default:
				if responseBuilder(update.Message.Text) != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка.\nОтправь название страны")
					bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
					bot.Send(msg)
				}
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка.\nОтправь название страны")
			bot.Send(msg)
		}
	}
}
