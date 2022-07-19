package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_"`
	SID         string             `bson:"_sid" json:"_"`
	UserSID     string             `bson:"user_sid" json:"_"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Path        string             `bson:"path" json:"_"`
	Created_at  string             `bson:"created_at" json:"_"`
}
