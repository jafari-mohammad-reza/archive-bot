package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NoteModel struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty" `
	AuthorId      primitive.ObjectID   `json:"author_id,omitempty" bson:"author_id,omitempty"`
	AttachmentsId []primitive.ObjectID `json:"attachments_id,omitempty" bson:"attachments_id,omitempty"`
}
