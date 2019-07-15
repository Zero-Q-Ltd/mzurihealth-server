package gen

import (
	"context"
	"log"

	"github.com/kisinga/mzurihealth/db"
	"github.com/kisinga/mzurihealth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) AdminCategory() AdminCategoryResolver {
	return &adminCategoryResolver{r}
}
func (r *Resolver) AdminInvite() AdminInviteResolver {
	return &adminInviteResolver{r}
}
func (r *Resolver) HospAdmin() HospAdminResolver {
	return &hospAdminResolver{r}
}
func (r *Resolver) Hospital() HospitalResolver {
	return &hospitalResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type adminCategoryResolver struct{ *Resolver }

func (r *adminCategoryResolver) _id(ctx context.Context, obj *models.AdminCategory) (string, error) {
	panic("not implemented")
}

type adminInviteResolver struct{ *Resolver }

func (r *adminInviteResolver) _id(ctx context.Context, obj *models.AdminInvite) (string, error) {
	panic("not implemented")
}

type hospAdminResolver struct{ *Resolver }

func (r *hospAdminResolver) _id(ctx context.Context, obj *models.HospAdmin) (string, error) {
	panic("not implemented")
}

type hospitalResolver struct{ *Resolver }

func (r *hospitalResolver) _id(ctx context.Context, obj *models.Hospital) (string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePatient(ctx context.Context, input models.NewPatient) (*models.Patient, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateAdmin(ctx context.Context, input *models.NewHospAdmin) (*models.HospAdmin, error) {
	collectionName := "hospadmins"
	// insertResult, err := db.InserDocument("", collectionName, input)
	// log.Print(insertResult)
	// str := fmt.Sprintf("%v", insertResult.InsertedID)
	var admin *models.HospAdmin
	objID, _ := primitive.ObjectIDFromHex("5d27e20cfe68ab3cfc380d7d")
	result := db.QueryDocument(ctx, "", collectionName, bson.D{{"_id", objID}})
	if result.Err() != nil {
		log.Fatalln("Failed to create request log file:", result.Err())

	}
	err := result.Decode(admin)

	return admin, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Patient(ctx context.Context, id string) ([]*models.Patient, error) {
	panic("not implemented")
}
func (r *queryResolver) PatientNotes(ctx context.Context, patientid string, from *int, to *int) ([]*models.Patientnote, error) {
	panic("not implemented")
}
func (r *queryResolver) HospAdmin(ctx context.Context, id string) (*models.HospAdmin, error) {
	panic("not implemented")
}
func (r *queryResolver) AdminCategories(ctx context.Context) ([]*models.AdminCategory, error) {
	panic("not implemented")
}
func (r *queryResolver) Hospital(ctx context.Context, id string) (*models.Hospital, error) {
	panic("not implemented")
}
func (r *queryResolver) HospFile(ctx context.Context, id *string) (*models.HospFile, error) {
	panic("not implemented")
}
func (r *queryResolver) AllInsurance(ctx context.Context) ([]*models.Insurance, error) {
	panic("not implemented")
}
func (r *queryResolver) PaymentChannels(ctx context.Context, hospitalid string) ([]*models.PaymentChannel, error) {
	panic("not implemented")
}
func (r *queryResolver) AllPaymentMethods(ctx context.Context) ([]*models.PaymentMethod, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) QueueChanged(ctx context.Context, id string) (<-chan *models.Patient, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) UserChanged(ctx context.Context, id string) (<-chan *models.HospAdmin, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) HospChanged(ctx context.Context, id *string) (<-chan *models.Hospital, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) PatientChanged(ctx context.Context, id *string) (<-chan *models.Patient, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) ProcedureChanged(ctx context.Context, id *string) (<-chan *models.RawProcedure, error) {
	panic("not implemented")
}
