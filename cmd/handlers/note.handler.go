package handlers

import (
	"archive-bot/cmd/db"
	middleware "archive-bot/cmd/middlewares"
	"archive-bot/cmd/models"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

import (
	"context"
	"time"
)

const ctxTimeout = time.Second * 2

// func /notes - *show list of user notes by first 50 character of that note or caption of that attachment*
func GetNoteHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	userKey := update.Message.From.String()
	user, ok := middleware.AuthorizedUsers[userKey]
	if !ok {
		return errors.New("user not authorized")
	}
	noteRepo := db.GetNoteRepo()

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	notes, err := noteRepo.GetAllBy(ctx, "author_id", *user)
	if err != nil {
		return err
	}
	for _, note := range notes {
		msgText, inlineMarkup := getNoteMsg(note)
		if msgText == "" {
			continue
		}
		sendMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		sendMsg.ReplyMarkup = inlineMarkup
		_, err = bot.Send(sendMsg)
		if err != nil {
			return err
		}
	}
	return nil
}
func getNoteMsg(note models.NoteModel) (string, tgbotapi.InlineKeyboardMarkup) {
	var content string
	if note.Content == "" {
		return "", tgbotapi.InlineKeyboardMarkup{}
	}
	if len(note.Content) > 50 {
		content = note.Content[:50] + "..."
	} else {
		content = note.Content
	}
	btnSee := tgbotapi.NewInlineKeyboardButtonSwitch("See Note", fmt.Sprintf("/see %s", note.ID))
	row1 := tgbotapi.NewInlineKeyboardRow(btnSee)
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(row1)
	return fmt.Sprintf("%s", content), inlineKeyboard
}
