package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NoteModel struct {
	BaseModel
	AuthorId      primitive.ObjectID   `json:"author_id,omitempty" bson:"author_id,omitempty"`
	AttachmentsId []primitive.ObjectID `json:"attachments_id,omitempty" bson:"attachments_id,omitempty"`
}
