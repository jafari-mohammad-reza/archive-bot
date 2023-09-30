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

	username := update.Message.From.String()
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
	if AuthorizedUsers[update.Message.From.String()] != nil {
		return nil
	} else {
		AuthorizedUsers[update.Message.From.String()] = &existUser.ID
	}
	return nil
}
