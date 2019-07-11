package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var client *mongo.Client
var defaultdb string

// ConnectDB is the entry point when the server is started that ensures a successful connection to a mongodb instance
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

// GetCollection returns a collection reference that can be used for reading or writing to db
func GetCollection(database string, collection string) *mongo.Collection {
	if database == "" {
		return client.Database(defaultdb).Collection(collection)
	}
	return client.Database(database).Collection(collection)
}

//InsertManyDocuments is the universal function that inserts MANY docs to the same collection
func InsertManyDocuments(database string, collection string, documents []interface{}) (*mongo.InsertManyResult, error) {
	var ref *mongo.Collection
	if database == "" {
		ref = client.Database(defaultdb).Collection(collection)
	} else {
		ref = client.Database(database).Collection(collection)
	}
	return ref.InsertMany(context.TODO(), documents)
}

//InserDocument is the universal function that inserts ONE doc to a collection
func InserDocument(database string, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	var ref *mongo.Collection
	if database == "" {
		ref = client.Database(defaultdb).Collection(collection)
	} else {
		ref = client.Database(database).Collection(collection)
	}
	return ref.InsertOne(context.TODO(), document)
}

//QueryDocument sends a query to the db
//By design, even fetching a specific document is a query, so whether the id is known or not does not change the structure of the function
//It is ip to the calling function to decode the document
func QueryDocument(database string, collection string, query bson.M) *mongo.SingleResult {
	return GetCollection(database, collection).FindOne(context.TODO(), query)
}

//DeleteDocument removes a document from the database
func DeleteDocument(database string, collection string, id bson.ObjectId) (*mongo.DeleteResult, error) {
	return GetCollection(database, collection).DeleteOne(context.TODO(), id)
}

//DeleteDocuments removes many documents from the database
func DeleteDocuments(database string, collection string, id []bson.ObjectId) (*mongo.DeleteResult, error) {
	return GetCollection(database, collection).DeleteMany(context.TODO(), id)
}

// QueryAggregate reads a many documents from the database
//A set of many complext queries can be built and batched in the same read operation
//Be very careful though, as pooply indexed database can be a huge bottleneck
//It is ip to the calling function to decode the documents
func QueryAggregate(database string, collection string, query []bson.M) (*mongo.Cursor, error) {
	return client.Database(database).Collection(collection).Aggregate(context.TODO(), query)
}

//CloseSession is good cleanup code when the server exits
func CloseSession() error {
	return client.Disconnect(context.TODO())
}
