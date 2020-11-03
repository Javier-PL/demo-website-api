package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" `
	Title string             `json:"title" form:"title" bson:"title"`
}
