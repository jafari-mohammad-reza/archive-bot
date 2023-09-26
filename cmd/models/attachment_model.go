package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AttachmentModel struct {
	BaseModel
	AuthorId primitive.ObjectID `json:"author_id,omitempty" bson:"author_id,omitempty"`
}
