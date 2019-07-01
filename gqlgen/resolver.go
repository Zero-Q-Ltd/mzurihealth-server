package gqlgen

import (
	"context"

	"github.com/kisinga/mzurihealth-server/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Patient() PatientResolver {
	return &patientResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePatient(ctx context.Context, input NewPatient) (*models.Patient, error) {
	panic("not implemented")
}

type patientResolver struct{ *Resolver }

func (r *patientResolver) Personalinfo(ctx context.Context, obj *models.Patient) (string, error) {
	panic("not implemented")
}
func (r *patientResolver) Fileinfo(ctx context.Context, obj *models.Patient) (*HospFile, error) {
	panic("not implemented")
}
func (r *patientResolver) Done(ctx context.Context, obj *models.Patient) (bool, error) {
	panic("not implemented")
}
func (r *patientResolver) User(ctx context.Context, obj *models.Patient) (*User, error) {
	panic("not implemented")
}
func (r *patientResolver) Nextofkin(ctx context.Context, obj *models.Patient) (*Nextofkin, error) {
	panic("not implemented")
}
func (r *patientResolver) Insurance(ctx context.Context, obj *models.Patient) ([]*Insurance, error) {
	panic("not implemented")
}
func (r *patientResolver) Medicalinfo(ctx context.Context, obj *models.Patient) (*Medicalinfo, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Patient(ctx context.Context, id string) ([]*models.Patient, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) Queuechanged(ctx context.Context, id string) (<-chan *models.Patient, error) {
	panic("not implemented")
}
