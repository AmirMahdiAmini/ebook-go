package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_"`
	SID           string             `bson:"_sid" json:"_"`
	UserName      string             `bson:"username" json:"username" binding:"required"`
	Phone         string             `bson:"phone" json:"phone" binding:"required"`
	Password      string             `bson:"password" json:"password" binding:"required"`
	Is_verified   bool               `bson:"is_verified" json:"_"`
	Notifications []string           `bson:"notifications" json:"_"`
	Created_at    string             `bson:"created_at" json:"_"`
}
