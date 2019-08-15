package gen

import (
	"context"

	"github.com/kisinga/mzurihealth/models"
	"github.com/kisinga/mzurihealth/resolvers"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePatient(ctx context.Context, input models.NewPatient) (*models.Patient, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateAdmin(ctx context.Context, input *models.NewAdmin) (*models.HospAdmin, error) {
	return resolvers.CreateAdmin(ctx, input)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Patient(ctx context.Context, id string) ([]*models.Patient, error) {
	panic("not implemented")
}
func (r *queryResolver) PatientNotes(ctx context.Context, patientid string, from *int, to *int) ([]*models.PatientNote, error) {
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
