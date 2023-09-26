package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" `
	UserName    string             `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Notes       []NoteModel        `json:"notes,omitempty" bson:"notes,omitempty"`
	Attachments []AttachmentModel  `json:"attachments,omitempty" bson:"attachments,omitempty"`
}
