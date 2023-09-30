package handlers

import (
	"archive-bot/cmd/db"
	middleware "archive-bot/cmd/middlewares"
	"archive-bot/cmd/models"
	"context"
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func /save - *save user message as a note and if message contains any attachment save attachment as well*
func SaveHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {

	// Check if message contains any text beside the '/save' command
	text := getTextAfterSaveCommand(update)
	isText := text != ""
  ctx , cancel := context.WithTimeout(context.TODO() , time.Second*5)
  defer cancel()
  var err error
	if isText {
    err = 	saveTextNote(update, &text , ctx)
    if err != nil {
      return err
    }
	}

	// Check if message contains any document or photo
	isAttachment := update.Message.Document != nil || update.Message.Photo != nil
	if isAttachment {
		saveAttachment(update , ctx)
	}


	// If neither attachment nor text is found, send a message to the user
	if !isAttachment && !isText {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please provide something to save (attachment or text)")
		_, err = bot.Send(msg)
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
func saveTextNote(update *tgbotapi.Update, text *string , ctx context.Context)  error {
  noteRepo := db.GetNoteRepo()
  fmt.Println("NOTE REPO" , noteRepo)
  note := models.NoteModel{AuthorId:*middleware.AuthorizedUsers[update.Message.From.String()] ,Content: *text , ContentFormat: models.Text }
  _ , err := noteRepo.Create(ctx , note)
  if err != nil {
    return err
  }
	return    nil
}

func saveAttachment(update *tgbotapi.Update , ctx context.Context) (*[]*primitive.ObjectID , error) {
  attachmentRepo := db.GetAttachmentRepository()
	// documents := update.Message.Document
	// photos := update.Message.Photo
  authorId := middleware.AuthorizedUsers[update.Message.From.String()]
  attachment := models.AttachmentModel{AuthorId: *authorId}
  attachmentRepo.Create(ctx , attachment)
	return nil , nil
}
