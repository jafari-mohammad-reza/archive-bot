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
func ConnectToDB() (err error) {
	mongoUrl := os.Getenv("MONGO_URL")
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
  defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			disconnectErr := client.Disconnect(ctx)
			if disconnectErr != nil {
				err = disconnectErr
			}
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Println("Database connected successfully")
	}

	database := client.Database(os.Getenv("MONGO_DATABASE"))
	defineCollections(database)
	mongoClient = client
	mongoDatabase = database
  defineRepositories()
	return nil
}
func defineCollections(database *mongo.Database) {
	database.Collection("user")
	database.Collection("attachment")
	database.Collection("note")
}
func defineRepositories() {
  noteRepo = NewNoteRepository()
  attachmentRepo = NewAttachmentRepository()
}
func GetMongoClient() *mongo.Client {
	return mongoClient
}
func GetMongoDatabase() *mongo.Database {
	return mongoDatabase
}
