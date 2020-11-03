package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Photo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" `
	Title       string             `json:"title" form:"title" bson:"title"`
	Description string             `json:"description" form:"description" bson:"description"`
	Base64      string             `json:"base64" form:"base64" bson:"base64,omitempty"`
	Category    string             `json:"category" form:"category" bson:"category"`
	Type        string             `json:"type" form:"type" bson:"type"`
	CategoryID  string             `json:"categoryID" form:"categoryID" bson:"categoryID"`
}
