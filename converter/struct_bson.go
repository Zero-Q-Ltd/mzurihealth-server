package converter

import (
	"go.mongodb.org/mongo-driver/bson"
)

//StructToBson simply converts Stuct to BBON
func StructToBson(input interface{}) (bsondata bson.M) {
	// Long way but achieves the purpose
	// First convert to byte then to bson
	// I couldnt find a direct way of doing that without using a third party library Here: https://github.com/fatih/structs
	b, err := bson.Marshal(input)
	if err != nil {
		//@TODO impliment tracing
		panic("cant marshall struct to bson")
	}
	err = bson.Unmarshal([]byte(b), &bsondata)
	if err != nil {
		//@TODO impliment tracing
		panic("cant unmarshall struct to bson")
	}
	return
}
