package crud

// https://gist.github.com/radhakishans1378/852fd10a286a03bd76ffcad00aec417b

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	database "monkey/db"
	"monkey/models"
)

// ISSUES collection for all of our issues
const ISSUES = "issues"

func CreateOne(task models.Issue) error {
	// Get MongoDb
	collection, err := database.GetCollection(ISSUES)
	_, err = collection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}
	return nil
}

func CreateMany(list []models.Issue) error {

	// Map struct slice to interface slice as InsertMany accepts interface slice
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}

	// Get MongoDB
	collection, err := database.GetCollection(ISSUES)
	_, err = collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return err
	}

	return nil
}

func GetOneIssueByCode(code string) (models.Issue, error) {
	result := models.Issue{}
	filter := bson.D{primitive.E{Key: "code", Value: code}}
	collection, err := database.GetCollection(ISSUES)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetOneIssueById(id string) (models.Issue, error) {
	// ID must be converted to an ObjectID before it'll match anything...
	objectId, err := primitive.ObjectIDFromHex(id)
	result := models.Issue{}
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
	collection, err := database.GetCollection(ISSUES)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println(err)
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetAllIssues() ([]models.Issue, error) {
	filter := bson.D{{}} // All documents
	issues := []models.Issue{}

	collection, err := database.GetCollection(ISSUES)
	if err != nil {
		return issues, err
	}
	cur, findErr := collection.Find(context.TODO(), filter)
	if findErr != nil {
		return issues, findErr
	}

	for cur.Next(context.TODO()) {
		t := models.Issue{}
		err := cur.Decode(&t)

		if err != nil {
			// could result in a partial return...
			return issues, err
		}
		issues = append(issues, t)
	}

	if len(issues) == 0 {
		return issues, mongo.ErrNoDocuments
	}
	return issues, nil
}

func DeleteOne(code string) error {
	filter := bson.D{primitive.E{Key: "code", Value: code}}

	collection, err := database.GetCollection(ISSUES)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAll() error {
	selector := bson.D{{}}
	collection, err := database.GetCollection(ISSUES)
	_, err = collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return err
	}
	return nil
}
