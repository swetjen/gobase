package crud

// https://gist.github.com/radhakishans1378/852fd10a286a03bd76ffcad00aec417b
// TODO:  Migrate to bound methods on crud struct.

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	database "monkey/db"
	"time"
)

const TEST = "test"

type testItem struct {
	ID        primitive.ObjectID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func IsDatabaseWorking() error {
	item := testItem{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c, err := database.GetCollection(TEST)
	if err != nil {
		log.Fatalln(err)
	}

	r, insertErr := c.InsertOne(context.TODO(), item)

	if insertErr != nil {
		log.Fatalln(insertErr)
	}

	log.Printf("Test Documented Inserted (%s)\n", r.InsertedID)

	d, deleteErr := c.DeleteOne(context.TODO(), item)
	if deleteErr != nil {
		log.Fatalln(deleteErr)
	}
	log.Printf("Test Document Deleted (%v)\n", d.DeletedCount)
	return nil
}
