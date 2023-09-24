package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"strings"
)

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
	files, err := os.ReadDir("/home/yeezus/Desktop/projects/archive-bot/cmd/handlers")
	if err != nil {
		return nil, err
	}

	var names []string
	for _, file := range files {
		if !file.IsDir() {
			name := file.Name()
			parts := strings.Split(name, ".handler")
			if len(parts) > 0 {
				if parts[0] == "error" {
					continue
				}
				names = append(names, parts[0])
			}
		}
	}

	return &names, nil
}
