package cmd

import (
	"archive-bot/cmd/handlers"
	middleware "archive-bot/cmd/middlewares"
	"errors"
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
	authMiddleware := middleware.NewAuthMiddleware()

	for update := range updates {
		err := authMiddleware.Authorize(&update)
		if err != nil {
			// handle authorization error
			log.Println("Authorization failed:", err.Error())
			continue
		}

		var errCh chan error
		if update.Message.Text == "/start" {
			errCh = make(chan error, 1)
			go func() {
				errCh <- handlers.StartHandler(bot, update)
			}()
		} else {
			errCh = handleUpdate(bot, &update)
		}
		checkErr(errCh)
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) chan error {
	errCh := make(chan error, 1)

	go func() {
		messageTextList := strings.Split(update.Message.Text, "/")
		if len(messageTextList) < 2 {
			errCh <- handlers.InvalidCmdHandler(bot, update)
			return
		}

		commandText := messageTextList[1]
		switch commandText {
		case "contact":
			errCh <- handlers.ContactHandler(bot, update)
		case "help":
			errCh <- handlers.HelpHandler(bot, update)
		default:
			errCh <- handlers.InvalidCmdHandler(bot, update)
		}
	}()

	return errCh
}
func checkErr(errCh chan error) {
	err := <-errCh
	if err != nil {
		log.Fatal("error at handling commands", err.Error())
	}
}
