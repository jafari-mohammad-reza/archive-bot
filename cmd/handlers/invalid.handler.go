package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func InvalidCmdHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, "https://i.pinimg.com/736x/e0/0c/62/e00c628d83109a800c86a3725cf6a295.jpg")
	msg.Caption = fmt.Sprintf(`There is no command as "%s" use /help for list of commmands and their usage`, update.Message.Text)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending photo: %v", err)
		return err
	}
	return nil
}
