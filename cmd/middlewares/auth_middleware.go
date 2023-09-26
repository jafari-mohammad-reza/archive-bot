package middleware

import (
	"archive-bot/cmd/db/repo"
	"archive-bot/cmd/models"
	"context"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AuthMiddleware struct {
	userRepo *repo.UserRepository
}

func NewAuthMiddleware() *AuthMiddleware {
	userRepo := repo.NewUserRepository()
	return &AuthMiddleware{userRepo}
}

func (m *AuthMiddleware) Authorize(update *tgbotapi.Update) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	username := update.Message.From.String()
	fmt.Println("USERNAME", username)
	_, err := m.userRepo.GetBy(ctx, "user_name", username)
	if errors.Is(err, mongo.ErrNoDocuments) {
		user := models.UserModel{UserName: username}
		if err := m.userRepo.Create(ctx, user); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}
