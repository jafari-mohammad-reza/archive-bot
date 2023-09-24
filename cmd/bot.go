package cmd

import (
	"archive-bot/cmd/handlers"
	"errors"
	"fmt"
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
	var err error
	for update := range updates {
		if update.Message.Text == "/start" {
			err = handlers.StartHandler(bot, update)
		} else {
			handleUpdate(bot, &update)
		}
		if err != nil {
			fmt.Print("Error ", err.Error())
			handlers.ErrorHandler(bot, update)
		}
	}

}
func handleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	text := update.Message.Text
	text = strings.Split(text, "/")[1]
	switch text {
	case "contact":
		return handlers.ContactHandler(bot, update)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "invalid message")
		bot.Send(msg)
		return nil
	}

}
