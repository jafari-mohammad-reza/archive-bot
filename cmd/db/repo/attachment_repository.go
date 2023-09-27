package repo

import (
	"archive-bot/cmd/db"
	"archive-bot/cmd/models"
)

type AttachmentRepository struct {
	MongoDbAbstractRepository[models.AttachmentModel]
}

func NewAttachmentRepository() *AttachmentRepository {
	noteCollection := db.GetMongoDatabase().Collection("attachment")
	return &AttachmentRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[models.AttachmentModel]{
			Collection: noteCollection,
		},
	}
}
