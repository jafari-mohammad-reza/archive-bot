package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at;omitempty" json:"created_at"`
	ModifyAt  primitive.DateTime `bson:"modify_at;omitempty" json:"modify_at"`
}
