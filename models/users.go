package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Name           string             `json:"user" bson:"user"`
	Email          string             `json:"email" bson:"email"`
	Phone          string             `json:"phone" bson:"phone"`
	HashedPassword string             `json:"hashed_password" bson:"hashed_password"`
	IsActive       bool               `json:"is_active" bson:"is_active"`
	IsSuperUser    bool               `json:"is_superuser" bson:"is_superuser"`
}

type UserCreate struct {
	Name        string `json:"user" bson:"user"`
	Email       string `json:"email" bson:"email"`
	Password    string `json:"password"`
	IsActive    bool   `json:"is_active" bson:"is_active"`
	IsSuperUser bool   `json:"is_superuser" bson:"is_superuser"`
}
