package models

import "go.mongodb.org/mongo-driver/bson/primitive"
type NoteFormat string
var (
  Text NoteFormat = "TEXT"
  MarkDown NoteFormat = "MARKDOWN"
)
type NoteModel struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty" `
	AuthorId      primitive.ObjectID   `json:"author_id,omitempty" bson:"author_id,omitempty"`
  Content string  `json:"content,omitempty" bson:"content,omitempty"`
  ContentFormat NoteFormat `json:"content_format;omitempty" bson:"content_format;omitempty"`
	AttachmentsId []primitive.ObjectID `json:"attachments_id,omitempty" bson:"attachments_id,omitempty"`
}
