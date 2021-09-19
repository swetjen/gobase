package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Contact struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Message   string             `bson:"message" json:"message"`
	Email     string             `bson:"email" json:"email"`
}
