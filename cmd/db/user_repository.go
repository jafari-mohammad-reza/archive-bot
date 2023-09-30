package db

import (
	"archive-bot/cmd/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
	MongoDbAbstractRepository[models.UserModel]
}

func NewUserRepository() *UserRepository {

	userCollection := GetMongoDatabase().Collection("users")

	return &UserRepository{
		db: GetMongoDatabase(),
		MongoDbAbstractRepository: MongoDbAbstractRepository[models.UserModel]{
			Collection: userCollection,
		},
	}
}
