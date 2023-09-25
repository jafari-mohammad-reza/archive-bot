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
	if token == "" {
		return errors.New("token does not exist")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		return err
	}

	setupBotHandlers(bot, updates)

	return nil
}

func setupBotHandlers(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		var handlerError error
		if update.Message.Text == "/start" {
			handlerError = handlers.StartHandler(bot, update)
		} else {
			handlerError = handleUpdate(bot, &update)
		}
		if handlerError != nil {
			fmt.Print("Error ", handlerError.Error())
			handlers.ErrorHandler(bot, update)
		}
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	messageTextList := strings.Split(update.Message.Text, "/")
	if len(messageTextList) < 2 {
		return handlers.InvalidCmdHandler(bot, update)
	}
	commandText := messageTextList[1]
	switch commandText {
	case "contact":
		return handlers.ContactHandler(bot, update)
	case "help":
		return handlers.HelpHandler(bot, update)
	default:
		return handlers.InvalidCmdHandler(bot, update)
	}
}
