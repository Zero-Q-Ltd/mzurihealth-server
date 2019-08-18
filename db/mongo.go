package db

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kisinga/mzurihealth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var defaultdb string

// ConnectDB is the entry point when the server is started that ensures a successful connection to a mongodb instance
func ConnectDB(ctx context.Context, defaultdbstring string) (err error) {
	fmt.Println("Connecting to Db........")

	defaultdb = defaultdbstring
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientvar, err := mongo.Connect(ctx, clientOptions)
	client = clientvar

	if err != nil {
		// log.Error(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("Error connecting to Db........\n", err)
		panic("Error connecting to Db........")
	}
	fmt.Println("Connected to MongoDB!")
	return err
}

// GetCollection returns a collection reference that can be used for reading or writing to db
func GetCollection(ctx context.Context, database string, collection string) *mongo.Collection {
	if database == "" {
		return client.Database(defaultdb).Collection(collection)
	}
	return client.Database(database).Collection(collection)
}

//ValidateAndGetUser reads the user from the database
func ValidateAndGetUser(ctx context.Context, cookie http.Cookie) (admin models.HospAdmin, err error) {
	userID, _ := primitive.ObjectIDFromHex(cookie.Value)
	res := QueryDocument(ctx, "", "hospadmins", bson.D{{"_id", userID}})
	if res.Err() != nil {
		return admin, res.Err()
	}
	err = res.Decode(admin)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

//InserDocument is the universal function that inserts ONE doc to a collection
//Accepts an optional id field and automatically adds one to the
func InserDocument(ctx context.Context, database string, collection string, id string, document bson.M) (*mongo.InsertOneResult, error) {
	var ref *mongo.Collection
	if database == "" {
		ref = client.Database(defaultdb).Collection(collection)
	} else {
		ref = client.Database(database).Collection(collection)
	}
	/**
	Make Mongo Automatically create an ID
	**/
	fmt.Printf("%#v", document)
	if id == "" {
		delete(document, "_id")
	}
	insertres, err := ref.InsertOne(ctx, document)
	if err != nil {
		dbError(ctx, "InserDocument", document, err)
	}
	return insertres, err
}

//InsertManyDocuments is the universal function that inserts MANY docs to the same collection
func InsertManyDocuments(ctx context.Context, database string, collection string, documents []interface{}) (*mongo.InsertManyResult, error) {
	var ref *mongo.Collection
	if database == "" {
		ref = client.Database(defaultdb).Collection(collection)
	} else {
		ref = client.Database(database).Collection(collection)
	}
	insertres, err := ref.InsertMany(ctx, documents)
	if err != nil {
		dbError(ctx, "InsertManyDocuments", documents, err)
	}
	return insertres, err
}

// UpdateDocument modifies a doc given the id and data
func UpdateDocument(ctx context.Context, database string, collection string, id bson.M, document interface{}) (*mongo.UpdateResult, error) {
	res, err := GetCollection(ctx, database, collection).UpdateOne(ctx, id, document)
	if err != nil {
		dbError(ctx, "UpdateDocument", document, err)
	}
	return res, err
}

//DeleteDocument removes a document from the database
func DeleteDocument(ctx context.Context, database string, collection string, id bson.M) (*mongo.DeleteResult, error) {
	res, err := GetCollection(ctx, database, collection).DeleteOne(ctx, id)
	if err != nil {
		dbError(ctx, "DeleteDocument", id, err)
	}
	return res, err
}

//DeleteDocuments removes many documents from the database
func DeleteDocuments(ctx context.Context, database string, collection string, id []bson.M) (*mongo.DeleteResult, error) {
	res, err := GetCollection(ctx, database, collection).DeleteMany(ctx, id)
	if err != nil {
		dbError(ctx, "DeleteDocuments", id, err)
	}
	return res, err
}

//QueryDocument sends a query to the db
//By design, even fetching a specific document is a query, so whether the id is known or not does not change the structure of the function.
//It is Up to the calling function to decode the document
func QueryDocument(ctx context.Context, database string, collection string, query bson.D) *mongo.SingleResult {
	res := GetCollection(ctx, database, collection).FindOne(ctx, query)
	if res.Err() != nil {
		dbError(ctx, "QueryDocument", query, res.Err())
	}
	return res
}

// QueryAggregate reads a many documents from the database
//A set of many complext queries can be built and batched in the same read operation.
//Be very careful though, as pooply indexed database can be a huge bottleneck.
//It is up to the calling function to decode the documents
func QueryAggregate(ctx context.Context, database string, collection string, query []bson.M) (*mongo.Cursor, error) {
	res, err := client.Database(database).Collection(collection).Aggregate(ctx, query)
	if err != nil {
		dbError(ctx, "QueryAggregate", query, err)
	}
	return res, err
}

//This logs any error occured when performing any CRUD operaion to db
func dbError(ctx context.Context, function string, params interface{}, err error) {
	fmt.Print(("Error" + function + err.Error()))
}

//CloseSession is good cleanup code when the server exits
func CloseSession(ctx context.Context) {
	err := client.Disconnect(ctx)
	if err != nil {
		// log.Warn().Msg(err.Error())
	}
}
