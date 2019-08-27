package resolvers

import (
	"context"

	"github.com/kisinga/mzurihealth/converter"
	"github.com/kisinga/mzurihealth/db"
	"github.com/kisinga/mzurihealth/models"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateAdmin creates a new admin in the database given all basic params
func CreateAdmin(ctx context.Context, input *models.NewAdmin) (*models.HospAdmin, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateAdmin")
	defer span.Finish()
	collectionName := "hospadmins"
	// pass, err := bcrypt.GenerateFromPassword([]byte(input.Pass), bcrypt.DefaultCost)
	// if err != nil {
	// 	span.LogFields(
	// 		log.String("event", "soft error"),
	// 		log.String("type", "Error Encrypting"),
	// 		log.Error(err))
	// 	return nil, err
	// }
	insertResult, err := db.InserDocument(ctx, "", collectionName, "", converter.StructToBson(input))

	var admin *models.HospAdmin
	// admin.Authdata.Pass = string(pass)

	str, _ := insertResult.InsertedID.(primitive.ObjectID)
	objID, _ := primitive.ObjectIDFromHex(str.Hex())
	result := db.QueryDocument(ctx, "", collectionName, bson.D{{"_id", objID}})
	err = result.Decode(admin)
	if err != nil {
		span.LogFields(
			log.String("event", "soft error"),
			log.String("type", "Error Converting"),
			log.Error(err))
		return admin, err

	}
	return admin, err

}
