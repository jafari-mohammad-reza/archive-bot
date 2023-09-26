package repo

import (
	"archive-bot/cmd/db"
	"archive-bot/cmd/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
	MongoDbAbstractRepository[models.UserModel]
}

func NewUserRepository() *UserRepository {

	userCollection := db.GetMongoDatabase().Collection("users")

	return &UserRepository{
		db: db.GetMongoDatabase(),
		MongoDbAbstractRepository: MongoDbAbstractRepository[models.UserModel]{
			Collection: userCollection,
		},
	}
}
