package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/fs"
	"os"
	"strings"
	"sync"
)

var once sync.Once
var names []string

func StartHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	handlers, err := handlerNames()
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, `welcome`)

	for _, handler := range *handlers {
		msg.Text += "\n/" + handler
	}

	msg.ReplyToMessageID = update.Message.MessageID
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func handlerNames() (*[]string, error) {
	var err error // Declare here

	once.Do(func() {
		var files []fs.DirEntry // Declare here

		files, err = os.ReadDir("./cmd/handlers") // Assigns to the outer err
		if err != nil {
			return
		}

		for _, file := range files {
			if !file.IsDir() {
				name := file.Name()
				parts := strings.Split(name, ".handler")
				fmt.Println(parts)
				if len(parts) > 0 {
					if isInvalidName(parts[0]) {
						continue
					}
					names = append(names, parts[0])
				}
			}
		}
	})

	// Here err is accessible
	if err != nil {
		return nil, err
	}

	fmt.Println("names", names)
	return &names, nil
}
func isInvalidName(name string) bool {
	invalidNames := []string{"error", "invalid"}
	for _, n := range invalidNames {
		if name == n {
			return true
		}
	}
	return false
}
