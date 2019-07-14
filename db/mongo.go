package db

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		log.Error()
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	fmt.Println("Connected to MongoDB!")
	//dont forget to close the connection
}

// GetCollection returns a collection reference that can be used for reading or writing to db
func GetCollection(database string, collection string) *mongo.Collection {
	if database == "" {
		return client.Database(defaultdb).Collection(collection)
	}
	return client.Database(database).Collection(collection)
}

//InserDocument is the universal function that inserts ONE doc to a collection
func InserDocument(database string, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	var ref *mongo.Collection
	if database == "" {
		ref = client.Database(defaultdb).Collection(collection)
	} else {
		ref = client.Database(database).Collection(collection)
	}
	insertres, err := ref.InsertOne(context.TODO(), document)
	if err != nil {
		dbError("InserDocument", document, err)
	}
	return insertres, err
}

//InsertManyDocuments is the universal function that inserts MANY docs to the same collection
func InsertManyDocuments(database string, collection string, documents []interface{}) (*mongo.InsertManyResult, error) {
	var ref *mongo.Collection
	if database == "" {
		ref = client.Database(defaultdb).Collection(collection)
	} else {
		ref = client.Database(database).Collection(collection)
	}
	insertres, err := ref.InsertMany(context.TODO(), documents)
	if err != nil {
		dbError("InsertManyDocuments", documents, err)
	}
	return insertres, err
}

// UpdateDocument modifies a doc given the id and data
func UpdateDocument(database string, collection string, id bson.M, document interface{}) (*mongo.UpdateResult, error) {
	res, err := GetCollection(database, collection).UpdateOne(context.TODO(), id, document)
	if err != nil {
		dbError("UpdateDocument", document, err)
	}
	return res, err
}

//DeleteDocument removes a document from the database
func DeleteDocument(database string, collection string, id bson.M) (*mongo.DeleteResult, error) {
	res, err := GetCollection(database, collection).DeleteOne(context.TODO(), id)
	if err != nil {
		dbError("DeleteDocument", id, err)
	}
	return res, err
}

//DeleteDocuments removes many documents from the database
func DeleteDocuments(database string, collection string, id []bson.M) (*mongo.DeleteResult, error) {
	res, err := GetCollection(database, collection).DeleteMany(context.TODO(), id)
	if err != nil {
		dbError("DeleteDocuments", id, err)
	}
	return res, err
}

//QueryDocument sends a query to the db
//By design, even fetching a specific document is a query, so whether the id is known or not does not change the structure of the function.
//It is Up to the calling function to decode the document
func QueryDocument(database string, collection string, query bson.D) *mongo.SingleResult {
	res := GetCollection(database, collection).FindOne(context.TODO(), query)
	// log.Info().Msg("Query", +query.)

	if res.Err() != nil {
		log.Warn().Msg(res.Err().Error())
		dbError("QueryDocument", query, res.Err())
	}
	return res
}

// QueryAggregate reads a many documents from the database
//A set of many complext queries can be built and batched in the same read operation.
//Be very careful though, as pooply indexed database can be a huge bottleneck.
//It is up to the calling function to decode the documents
func QueryAggregate(database string, collection string, query []bson.M) (*mongo.Cursor, error) {
	res, err := client.Database(database).Collection(collection).Aggregate(context.TODO(), query)
	if err != nil {
		dbError("QueryAggregate", query, err)
	}
	return res, err
}

//This logs any error occured when performing any CRUD operaion to db
func dbError(function string, params interface{}, err error) {
	log.Warn().Msg("Error" + function + err.Error())
}

//CloseSession is good cleanup code when the server exits
func CloseSession() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Warn().Msg(err.Error())
	}
}
