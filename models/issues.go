package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Issue struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	Title       string             `bson:"title" json:"title"`
	Code        string             `bson:"code" json:"code"`
	Description string             `bson:"description" json:"description"`
	Completed   bool               `bson:"completed" json:"completed"`
}
