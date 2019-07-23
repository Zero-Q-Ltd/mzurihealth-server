package converter

import (
	"go.mongodb.org/mongo-driver/bson"
)

//StructToBson simply converts Stuct to BBON
func StructToBson(input interface{}) (bsondata bson.M) {
	b, err := bson.Marshal(input)
	if err != nil {
		//@TODO impliment tracing
		panic("cant convert struct to bson")
	}
	bson.Unmarshal([]byte(b), &bsondata)
	return
}
