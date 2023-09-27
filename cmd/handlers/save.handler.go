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

	// Check if message contains any text beside the '/save' command
	text := getTextAfterSaveCommand(update)
	isText := text != ""

	if isText {
		saveTextNote(update, &text)
	}

	// If neither attachment nor text is found, send a message to the user
	if !isAttachment && !isText {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please provide something to save (attachment or text)")
		_, err := bot.Send(msg)
		return err
	}

	return nil
}

// getTextAfterSaveCommand extracts text after '/save' command from an update message or caption
func getTextAfterSaveCommand(update *tgbotapi.Update) string {
	var text string
	possibleTexts := []string{update.Message.Text, update.Message.Caption}

	for _, possibleText := range possibleTexts {
		splitText := strings.Split(possibleText, "/save")
		if len(splitText) >= 2 {
			// Remove leading and trailing whitespaces
			text = strings.TrimSpace(splitText[1])
			if text != "" {
				break
			}
		}
	}

	return text
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
