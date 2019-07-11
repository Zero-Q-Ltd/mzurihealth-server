package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var defaultdb string

func ConnectDB(defaultdbstring string) {
	fmt.Println("Connecting to Db........")

	defaultdb = defaultdbstring
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientvar, err := mongo.Connect(context.TODO(), clientOptions)
	client = clientvar
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

}
func GetCollection(database string, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}

// func GetDocument(database string, collection string) *mongo.Collection {
// 	return client.Database(database).Collection(collection).
// }

func CloseSession() error {
	return client.Disconnect(context.TODO())
}
