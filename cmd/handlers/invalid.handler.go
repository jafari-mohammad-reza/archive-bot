package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func InvalidCmdHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	var msg *tgbotapi.Message
	if update.Message != nil {
		msg = update.Message
	} else if update.CallbackQuery != nil {
		msg = update.CallbackQuery.Message
	}
	sendMsg := tgbotapi.NewPhotoShare(msg.Chat.ID, "https://i.pinimg.com/736x/e0/0c/62/e00c628d83109a800c86a3725cf6a295.jpg")

	sendMsg.Caption = fmt.Sprintf(`There is no command as "%s" use /help for list of commmands and their usage`, msg.Text)
	_, err := bot.Send(sendMsg)
	return err
}
