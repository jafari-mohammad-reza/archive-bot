package cmd

import (
	"archive-bot/cmd/handlers"
	middleware "archive-bot/cmd/middlewares"
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

	// bot.Debug = true

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
	limiter := middleware.NewRateLimiter(10)

	for update := range updates {
		errCh := handleUpdateWithMiddleware(bot, update, authMiddleware, limiter)
		checkErr(errCh)
	}
}

func handleUpdateWithMiddleware(bot *tgbotapi.BotAPI, update tgbotapi.Update, authMiddleware *middleware.AuthMiddleware, limiter *middleware.RateLimiter) chan error {
	errChan := make(chan error)

	go func() {
		if err := authenticateAndUpdateRate(authMiddleware, limiter, &update); err != nil {
			errChan <- err
			return
		}

		if update.Message.Text == "/start" {
			errChan <- <-startHandlerAsync(bot, &update)
		} else {
			errChan <- <-handleUpdate(bot, &update)
		}
	}()

	return errChan
}

func authenticateAndUpdateRate(authMiddleware *middleware.AuthMiddleware, limiter *middleware.RateLimiter, update *tgbotapi.Update) error {
	err := authMiddleware.Authorize(update)
	if handleError("Authorization failed:", err) {
		return err
	}

	err = limiter.Request(update.Message.From.String())
	if handleError("You're doing that too often. Please wait.", err) {
		return err
	}

	return nil
}

func handleError(message string, err error) bool {
	if err != nil {
		log.Println(message, err.Error())
		return true
	}
	return false
}

func startHandlerAsync(bot *tgbotapi.BotAPI, update *tgbotapi.Update) chan error {
	errCh := make(chan error, 1)
	go func() {
		errCh <- handlers.StartHandler(bot, *update)
	}()

	return errCh
}

// This function takes a bot and an update as input, and returns a channel of errors
func getCommand(update *tgbotapi.Update) (string, error) {
	if !strings.ContainsAny(update.Message.Text, "/") {
		return "", fmt.Errorf("invalid command")
	}

	messageTextList := strings.Split(update.Message.Text, "/")
	if len(messageTextList) < 2 {
		return "", fmt.Errorf("invalid command")
	}

	return strings.Split(messageTextList[1], " ")[0], nil
}

func handleMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update, commandText string) error {
	switch commandText {
	case "contact":
		return handlers.ContactHandler(bot, update)
	case "help":
		return handlers.HelpHandler(bot, update)
	case "save":
		return handlers.SaveHandler(bot, update)
	case "notes":
		return handlers.GetNoteHandler(bot, update)
	default:
		return handlers.InvalidCmdHandler(bot, update)
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) chan error {
	errCh := make(chan error, 1)

	go func() {
		// If the update contains a photo or document, handle it with the SaveHandler
		if update.Message.Text == "" && (update.Message.Photo != nil || update.Message.Document != nil) {
			errCh <- handlers.SaveHandler(bot, update)
			return
		}

		commandText, err := getCommand(update)

		if err != nil {
			errCh <- handlers.InvalidCmdHandler(bot, update)
			return
		}

		errCh <- handleMessage(bot, update, commandText)
	}()

	return errCh
}
func checkErr(errCh chan error) {
	err := <-errCh
	if err != nil {
		log.Fatal("error at handling commands", err.Error())
	}
}
