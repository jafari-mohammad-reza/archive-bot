package cmd

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strings"
)

func SetupBot() error {
	token := os.Getenv("TOKEN")
	if len(token) <= 0 {
		return errors.New("token not exist")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return nil
	}
	setupHandlers(bot, updates)
	return nil
}
func setupHandlers(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, `welcome 
use  /hello for more messages
use /help  for list of commands and their usages
use /contact for contacting to me
`)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		} else {
			handleUpdate(bot, &update)
		}
	}
}
func handleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	text := update.Message.Text
	text = strings.Split(text, "/")[1]
	switch text {
	case "hello":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hey")
		bot.Send(msg)
	case "contact":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "contact me by mohammadrexzajafari.dev@gmail.com or https://t.me/DaYeezus in telegram")
		bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "invalid message")
		bot.Send(msg)
	}

}
