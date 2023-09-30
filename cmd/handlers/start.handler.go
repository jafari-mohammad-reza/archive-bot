package handlers

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var once sync.Once
var names []string
// func /start  - *Starts the bot and provides a welcome message*.
func StartHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	handlers, err := handlerNames()
	if err != nil {
		return fmt.Errorf("unable to retrieve handler names: %v", err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, `welcome`)

	for _, handler := range *handlers {
		msg.Text += "\n/" + handler
	}

	msg.ReplyToMessageID = update.Message.MessageID
	if _, err = bot.Send(msg); err != nil {
		return fmt.Errorf("unable to send message: %v", err)
	}
	return nil
}

func handlerNames() (*[]string, error) {
	var readDirErr error

	once.Do(func() {
		var files []fs.DirEntry

		files, readDirErr = os.ReadDir("./cmd/handlers")
		if readDirErr != nil {
			return
		}

		for _, file := range files {
			if !file.IsDir() {
				name := strings.TrimSuffix(file.Name(), ".handler.go")
				if isInvalidName(name) {
					continue
				}
				names = append(names, name)
			}
		}
	})

	// Here err is accessible
	if readDirErr != nil {
		return nil, fmt.Errorf("error reading directory: %v", readDirErr)
	}

	return &names, nil
}

func isInvalidName(name string) bool {
	return name == "error" || name == "invalid"
}
