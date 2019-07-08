package gen

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) HospAdmin() HospAdminResolver {
	return &hospAdminResolver{r}
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

type hospAdminResolver struct{ *Resolver }

func (r *hospAdminResolver) _id(ctx context.Context, obj *HospAdmin) (string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePatient(ctx context.Context, input NewPatient) (*Patient, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Patient(ctx context.Context, id string) ([]*Patient, error) {
	panic("not implemented")
}
func (r *queryResolver) HospAdmin(ctx context.Context, id string) ([]*HospAdmin, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) QueueChanged(ctx context.Context, id string) (<-chan *Patient, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) UserChanged(ctx context.Context, id string) (<-chan *HospAdmin, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) HospChanged(ctx context.Context, id *string) (<-chan *Hospital, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) PatientChanged(ctx context.Context, id *string) (<-chan *Patient, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) ProcedureChanged(ctx context.Context, id *string) (<-chan *RawProcedure, error) {
	panic("not implemented")
}
