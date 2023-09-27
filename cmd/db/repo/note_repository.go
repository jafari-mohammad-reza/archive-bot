package repo

import (
	"archive-bot/cmd/db"
	"archive-bot/cmd/models"
)

type NoteRepository struct {
	MongoDbAbstractRepository[models.NoteModel]
}

func NewNoteRepository() *NoteRepository {
	noteCollection := db.GetMongoDatabase().Collection("note")
	return &NoteRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[models.NoteModel]{
			Collection: noteCollection,
		},
	}
}
