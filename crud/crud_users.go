package crud

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"monkey/core"
	database "monkey/db"
	"monkey/models"
)

// https://gist.github.com/radhakishans1378/852fd10a286a03bd76ffcad00aec417b

const (
	USERS = "users"
)

func GetAllUsers() ([]models.User, error) {
	users := []models.User{}
	filter := bson.D{{}}

	collection, err := database.GetCollection(USERS)

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return users, err
	}

	for cur.Next(context.TODO()) {
		each := models.User{}
		err := cur.Decode(&each)

		if err != nil {
			return users, err
		}
		users = append(users, each)

	}

	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}
	return users, nil
}

func GetUserByEmail(email string) (models.User, error) {
	user := models.User{}
	filter := bson.D{primitive.E{Key: "email", Value: email}}

	collection, err := database.GetCollection(USERS)
	err = collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateOneUser(user models.UserCreate) error {
	collection, err := database.GetCollection(USERS)
	if err != nil {
		return err
	}
	var userToDb = models.User{
		Name:           user.Name,
		Email:          user.Email,
		HashedPassword: core.HashPassword(user.Password),
		IsSuperUser:    true,
		IsActive:       true,
	}

	_, err = collection.InsertOne(context.TODO(), userToDb)
	if err != nil {
		return err
	}
	return nil
}

// Todo: Refactor to a generic Intialize DB that scaffolds all required tables and data

func InitializeUsers() {
	users, _ := GetAllUsers()
	if len(users) == 0 {
		log.Printf("<database=%s, UsersCollection=%s, UsersLength=%v>\tNo users in DB.  Adding first user...", core.Settings.Db, USERS, len(users))
		// Add our first user
		var firstUser = models.UserCreate{
			Email:    core.Settings.FirstSuperUser,
			Password: core.Settings.FirstSuperUserPassword,
		}
		err := CreateOneUser(firstUser)
		if err != nil {
			panic(err)
		}

	}
}
