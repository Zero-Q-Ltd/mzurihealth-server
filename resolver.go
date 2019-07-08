package mzuriheallth

import (
	"context"

	"github.com/kisinga/mzurihealth/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Admin() AdminResolver {
	return &adminResolver{r}
}
func (r *Resolver) HospFile() HospFileResolver {
	return &hospFileResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Patient() PatientResolver {
	return &patientResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) RawProcedure() RawProcedureResolver {
	return &rawProcedureResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type adminResolver struct{ *Resolver }

func (r *adminResolver) _id(ctx context.Context, obj *Admin) (string, error) {
	panic("not implemented")
}

type hospFileResolver struct{ *Resolver }

func (r *hospFileResolver) ID(ctx context.Context, obj *models.HospFile) (string, error) {
	panic("not implemented")
}
func (r *hospFileResolver) Date(ctx context.Context, obj *models.HospFile) (int, error) {
	panic("not implemented")
}
func (r *hospFileResolver) Lastvisit(ctx context.Context, obj *models.HospFile) (int, error) {
	panic("not implemented")
}
func (r *hospFileResolver) No(ctx context.Context, obj *models.HospFile) (string, error) {
	panic("not implemented")
}
func (r *hospFileResolver) Idno(ctx context.Context, obj *models.HospFile) (*string, error) {
	panic("not implemented")
}
func (r *hospFileResolver) Visitcount(ctx context.Context, obj *models.HospFile) (*int, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePatient(ctx context.Context, input NewPatient) (*models.Patient, error) {
	panic("not implemented")
}

type patientResolver struct{ *Resolver }

func (r *patientResolver) Personalinfo(ctx context.Context, obj *models.Patient) (string, error) {
	panic("not implemented")
}
func (r *patientResolver) Fileinfo(ctx context.Context, obj *models.Patient) (*models.HospFile, error) {
	panic("not implemented")
}
func (r *patientResolver) Done(ctx context.Context, obj *models.Patient) (bool, error) {
	panic("not implemented")
}
func (r *patientResolver) Admin(ctx context.Context, obj *models.Patient) (*Admin, error) {
	panic("not implemented")
}
func (r *patientResolver) ID(ctx context.Context, obj *models.Patient) (string, error) {
	panic("not implemented")
}
func (r *patientResolver) Parentid(ctx context.Context, obj *models.Patient) (*string, error) {
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
func (r *queryResolver) Admin(ctx context.Context, id string) ([]*Admin, error) {
	panic("not implemented")
}

type rawProcedureResolver struct{ *Resolver }

func (r *rawProcedureResolver) Name(ctx context.Context, obj *models.RawProcedure) (string, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) QueueChanged(ctx context.Context, id string) (<-chan *models.Patient, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) UserChanged(ctx context.Context, id string) (<-chan *Admin, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) HospChanged(ctx context.Context, id *string) (<-chan *Hosp, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) PatientChanged(ctx context.Context, id *string) (<-chan *models.Patient, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) ProcedureChanged(ctx context.Context, id *string) (<-chan *models.RawProcedure, error) {
	panic("not implemented")
}
