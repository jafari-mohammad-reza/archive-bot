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
		err := <-errCh
		if err != nil {
			log.Println("An error occured:", err.Error())
		}
	}
}

func handleMessage(update *tgbotapi.Update) string {
	if update.Message != nil {
		return update.Message.Text
	} else if update.InlineQuery != nil {
		return update.InlineQuery.Query
	} else if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Text
	}
	return ""
}

func handleUpdateWithMiddleware(bot *tgbotapi.BotAPI, update tgbotapi.Update, authMiddleware *middleware.AuthMiddleware, limiter *middleware.RateLimiter) chan error {
	errChan := make(chan error)

	go func() {
		if err := authenticateAndUpdateRate(authMiddleware, limiter, &update); err != nil {
			errChan <- err
			return
		}

		if handleMessage(&update) == "/start" {
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

	username := getUserName(update)
	err = limiter.Request(username)
	if handleError("Rate limit error:", err) {
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

func getUserName(update *tgbotapi.Update) string {
	var username string
	if update.Message != nil {
		username = update.Message.From.String()
	} else if update.InlineQuery != nil {
		username = update.InlineQuery.From.String()
	} else if update.CallbackQuery != nil {
		username = update.CallbackQuery.From.String()
	}
	return username
}

func handleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) chan error {
	errCh := make(chan error, 1)

	go func() {
		var messageText string
		if update.Message != nil {
			messageText = update.Message.Text
		} else if update.InlineQuery != nil {
			messageText = update.InlineQuery.Query
		} else if update.CallbackQuery != nil {
			messageText = update.CallbackQuery.Message.Text
		}

		if messageText == "" && (update.Message != nil && (update.Message.Photo != nil || update.Message.Document != nil)) {
			errCh <- handlers.SaveHandler(bot, update)
			return
		}

		messageTextList := strings.Split(messageText, "/")
		fmt.Println("messageTextList", messageTextList)
		if len(messageTextList) < 2 {
			errCh <- errors.New("invalid command")
			return
		}

		command := strings.Split(messageTextList[1], " ")[0]
		errCh <- handleMessageWithCommand(bot, update, command)
	}()

	return errCh
}

func handleMessageWithCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update, command string) error {
	fmt.Println("COMMAND", command)
	switch command {
	case "contact":
		return handlers.ContactHandler(bot, update)
	case "help":
		return handlers.HelpHandler(bot, update)
	case "save":
		return handlers.SaveHandler(bot, update)
	case "notes":
		return handlers.GetNoteHandler(bot, update)
	case "see":
		return handlers.GetNoteHandler(bot, update)
	default:
		return handlers.InvalidCmdHandler(bot, update)
	}
}
