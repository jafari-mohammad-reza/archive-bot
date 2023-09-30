package db

import (
	"archive-bot/cmd/models"
)
var noteRepo *NoteRepository
type NoteRepository struct {
	MongoDbAbstractRepository[models.NoteModel]
}

func NewNoteRepository() *NoteRepository {
	noteCollection := GetMongoDatabase().Collection("note")
	return &NoteRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[models.NoteModel]{
			Collection: noteCollection,
		},
	}
}

func GetNoteRepo() *NoteRepository {
  return noteRepo
}
