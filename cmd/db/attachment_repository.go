package db

import (
	"archive-bot/cmd/models"
)
var attachmentRepo *AttachmentRepository
type AttachmentRepository struct {
	MongoDbAbstractRepository[models.AttachmentModel]
}

func NewAttachmentRepository() *AttachmentRepository {
	noteCollection := GetMongoDatabase().Collection("attachment")
	return &AttachmentRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[models.AttachmentModel]{
			Collection: noteCollection,
		},
	}
}
func GetAttachmentRepository() *AttachmentRepository {
  return attachmentRepo
}
