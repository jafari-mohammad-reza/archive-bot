package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

// func /save - *save user message as a note and if message contains any attachment save attachment as well*
func SaveHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	// Check if message contains any document or photo
	isAttachment := update.Message.Document != nil || update.Message.Photo != nil
	if isAttachment {
		saveAttachment(update)
	}
	// check if message contains any text beside the command
	var text []string
	text = strings.Split(update.Message.Text, "/save")
	if !(len(text) >= 2) {
		text = strings.Split(update.Message.Caption, "/save")
	}

	if len(text) >= 2 {
		trimText := strings.ReplaceAll(text[1], " ", "")
		isText := trimText != ""
		if isText {
			saveTextNote(update, &trimText)
		}
	}

	return nil
}

func saveTextNote(update *tgbotapi.Update, text *string) error {
	fmt.Println("save text")
	return nil
}

func saveAttachment(update *tgbotapi.Update) error {
	fmt.Println("save attachment")
	//documents := update.Message.Document
	//photos := update.Message.Photo
	return nil
}
