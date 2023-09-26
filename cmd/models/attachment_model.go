package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AttachmentModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" `
	AuthorId primitive.ObjectID `json:"author_id,omitempty" bson:"author_id,omitempty"`
}
