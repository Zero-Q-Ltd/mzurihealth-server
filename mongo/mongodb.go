package mongo

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2"
)

func ConnectDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	err = client.Ping(ctx, readpref.Primary())

}
func GetCollection(collection string) *mgo.Collection { return db.C(collection) }

func CloseSession() { session.Close() }
