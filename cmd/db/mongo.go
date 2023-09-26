package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var mongoClient *mongo.Client
var mongoDatabase *mongo.Database

func ConnectToDB() error {
	mongoUrl := os.Getenv("MONGO_URL")
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*5)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Println("Database connected successfully")
	}
	if err != nil {
		return err
		defer func(client *mongo.Client, ctx context.Context) error {
			err := client.Disconnect(ctx)
			if err != nil {
				return err
			}
			return nil
		}(client, ctx)
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
