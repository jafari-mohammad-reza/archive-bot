package middleware

import (
	"archive-bot/cmd/db"
	"archive-bot/cmd/models"
	"context"
	"errors"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMiddleware struct {
	userRepo *db.UserRepository
}

var AuthorizedUsers map[string]*primitive.ObjectID

func NewAuthMiddleware() *AuthMiddleware {
	userRepo := db.NewUserRepository()
	AuthorizedUsers = make(map[string]*primitive.ObjectID)
	return &AuthMiddleware{userRepo}
}

func (m *AuthMiddleware) Authorize(update *tgbotapi.Update) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var username string
	if update.Message != nil {
		username = update.Message.From.String()
	} else if update.InlineQuery != nil {
		username = update.InlineQuery.From.String()
	} else if update.CallbackQuery != nil {
		username = update.CallbackQuery.From.String()
	} else {
		return errors.New("invalid user")
	}
	existUser, err := m.userRepo.GetBy(ctx, "user_name", username)
	if errors.Is(err, mongo.ErrNoDocuments) {
		user := models.UserModel{UserName: username}
		_, err := m.userRepo.Create(ctx, user)

		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	if AuthorizedUsers[username] != nil {
		return nil
	} else {
		AuthorizedUsers[username] = &existUser.ID
	}
	return nil
}
