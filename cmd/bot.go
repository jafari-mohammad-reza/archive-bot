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
	limiter := middleware.NewRateLimiter(10)
	for update := range updates {
		err := authMiddleware.Authorize(&update)
		if err != nil {
			log.Println("Authorization failed:", err.Error())
			continue
		}
		if err := limiter.Request(update.Message.From.String()); err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You're doing that too often. Please wait.")
			bot.Send(msg)
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

// This function takes a bot and an update as input, and returns a channel of errors
func handleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) chan error {
	// Create an error channel with a buffer size of 1
	errCh := make(chan error, 1)

	// Start a new goroutine to handle the update asynchronously
	go func() {
		// If the update contains a photo or document, handle it with the SaveHandler
		// and return immediately
		if update.Message.Text == "" && (update.Message.Photo != nil || update.Message.Document != nil) {
			errCh <- handlers.SaveHandler(bot, update)
			return
		}

		// If the update message doesn't contain any "/", treat it as an invalid command
		if !strings.ContainsAny(update.Message.Text, "/") {
			errCh <- handlers.InvalidCmdHandler(bot, update)
			return
		}

		// Split the update message by "/", and treat it as an invalid command if it's too short
		messageTextList := strings.Split(update.Message.Text, "/")
		if len(messageTextList) < 2 {
			errCh <- handlers.InvalidCmdHandler(bot, update)
			return
		}

		// Extract the command text and remove any trailing space
		commandText := strings.Split(messageTextList[1], " ")[0]

		// Dispatch the command to the appropriate handler
		switch commandText {
		case "contact":
			errCh <- handlers.ContactHandler(bot, update)
		case "help":
			errCh <- handlers.HelpHandler(bot, update)
		case "save":
			errCh <- handlers.SaveHandler(bot, update)
		default:
			// If the command is not recognized, treat it as an invalid command
			errCh <- handlers.InvalidCmdHandler(bot, update)
		}
	}()

	// Return the error channel so the caller can wait for the operation to finish and check for errors
	return errCh
}
func checkErr(errCh chan error) {
	err := <-errCh
	if err != nil {
		log.Fatal("error at handling commands", err.Error())
	}
}
