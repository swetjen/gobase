package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"monkey/core"
	"sync"
)

// https://levelup.gitconnected.com/working-with-mongodb-using-golang-754ead0c10c
// https://gist.github.com/radhakishans1378/a90eaf230e6409f239a81dc0b2246fd0

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

// GetMongoClient returns a mongo client
func GetMongoClient() (*mongo.Client, error) {
	// Setup Phase that's executed only once
	mongoOnce.Do(func() {
		fmt.Println(core.Settings.MongoUri)
		clientOptions := options.Client().ApplyURI(core.Settings.MongoUri)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
		clientInstance = client
	})
	// Returned on every call.
	return clientInstance, clientInstanceError
}

// GetCollection returns a Mongo Collection for provided Collection Name
func GetCollection(colName string) (*mongo.Collection, error) {
	client, err := GetMongoClient()
	collection := client.Database(core.Settings.Db).Collection(colName)
	if err != nil {
		return collection, err
	}
	return collection, nil
}
