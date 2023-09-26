package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var mongoClient *mongo.Client
var mongoDatabase *mongo.Database

func ConnectToDB() error {
	mongoUrl := os.Getenv("mongodb://mongo:27017")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return err
		defer func(client *mongo.Client, ctx context.Context) error {
			err := client.Disconnect(ctx)
			if err != nil {
				return err
			}
			return nil
		}(client, context.Background())
	}
	database := client.Database("archive-bot")
	mongoClient = client
	mongoDatabase = database
	return nil
}
func GetMongoClient() *mongo.Client {
	return mongoClient
}
func GetMongoDatabase() *mongo.Database {
	return mongoDatabase
}
func stringToObjectId(id string) (primitive.ObjectID, error) {
	obi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return obi, nil
}
