package crud

// https://gist.github.com/radhakishans1378/852fd10a286a03bd76ffcad00aec417b

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	database "monkey/db"
	"monkey/models"
)

const CONTACTS = "contacts"

func ContactsCreateOne(contact models.Contact) error {
	collection, err := database.GetCollection(CONTACTS)
	_, err = collection.InsertOne(context.TODO(), contact)
	if err != nil {
		return err
	}
	return nil
}

func ContactsGetAll() ([]models.Contact, error) {
	filter := bson.D{{}}
	contacts := []models.Contact{}

	collection, err := database.GetCollection(CONTACTS)
	if err != nil {
		return contacts, err
	}

	cur, findErr := collection.Find(context.TODO(), filter)
	if findErr != nil {
		return contacts, findErr
	}

	for cur.Next(context.TODO()) {
		t := models.Contact{}
		err := cur.Decode(&t)

		if err != nil {
			return contacts, err
		}

		contacts = append(contacts, t)
	}

	if len(contacts) == 0 {
		return contacts, mongo.ErrNoDocuments
	}
	return contacts, nil
}
