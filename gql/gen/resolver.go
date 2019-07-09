package gen

import (
	"context"

	"github.com/kisinga/mzurihealth/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

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

type queryResolver struct{ *Resolver }

func (r *queryResolver) Patient(ctx context.Context, id string) ([]*models.Patient, error) {
	panic("not implemented")
}
func (r *queryResolver) HospAdmin(ctx context.Context, id string) ([]*models.HospAdmin, error) {
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
