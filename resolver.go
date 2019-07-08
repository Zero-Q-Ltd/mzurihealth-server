package mzuriheallth

import (
	"context"

	"github.com/kisinga/mzurihealth/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Allergy() AllergyResolver {
	return &allergyResolver{r}
}
func (r *Resolver) Condition() ConditionResolver {
	return &conditionResolver{r}
}
func (r *Resolver) HospAdmin() HospAdminResolver {
	return &hospAdminResolver{r}
}
func (r *Resolver) HospFile() HospFileResolver {
	return &hospFileResolver{r}
}
func (r *Resolver) Hospital() HospitalResolver {
	return &hospitalResolver{r}
}
func (r *Resolver) Insurance() InsuranceResolver {
	return &insuranceResolver{r}
}
func (r *Resolver) Medicalinfo() MedicalinfoResolver {
	return &medicalinfoResolver{r}
}
func (r *Resolver) Metadata() MetadataResolver {
	return &metadataResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Nextofkin() NextofkinResolver {
	return &nextofkinResolver{r}
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
func (r *Resolver) Vitals() VitalsResolver {
	return &vitalsResolver{r}
}

type allergyResolver struct{ *Resolver }

func (r *allergyResolver) Allergytype(ctx context.Context, obj *models.Allergy) (*string, error) {
	panic("not implemented")
}
func (r *allergyResolver) Detail(ctx context.Context, obj *models.Allergy) (*string, error) {
	panic("not implemented")
}
func (r *allergyResolver) Metadata(ctx context.Context, obj *models.Allergy) (*models.Metadata, error) {
	panic("not implemented")
}

type conditionResolver struct{ *Resolver }

func (r *conditionResolver) Conditiontype(ctx context.Context, obj *models.Condition) (*string, error) {
	panic("not implemented")
}
func (r *conditionResolver) Detail(ctx context.Context, obj *models.Condition) (*string, error) {
	panic("not implemented")
}
func (r *conditionResolver) Metadata(ctx context.Context, obj *models.Condition) (*models.Metadata, error) {
	panic("not implemented")
}

type hospAdminResolver struct{ *Resolver }

func (r *hospAdminResolver) _id(ctx context.Context, obj *models.HospAdmin) (string, error) {
	panic("not implemented")
}
func (r *hospAdminResolver) Name(ctx context.Context, obj *models.HospAdmin) (string, error) {
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

type hospitalResolver struct{ *Resolver }

func (r *hospitalResolver) ID(ctx context.Context, obj *models.Hospital) (string, error) {
	panic("not implemented")
}
func (r *hospitalResolver) Location(ctx context.Context, obj *models.Hospital) (*Location, error) {
	panic("not implemented")
}
func (r *hospitalResolver) Name(ctx context.Context, obj *models.Hospital) (string, error) {
	panic("not implemented")
}

type insuranceResolver struct{ *Resolver }

func (r *insuranceResolver) ID(ctx context.Context, obj *models.Insurance) (string, error) {
	panic("not implemented")
}
func (r *insuranceResolver) Insuranceno(ctx context.Context, obj *models.Insurance) (*string, error) {
	panic("not implemented")
}

type medicalinfoResolver struct{ *Resolver }

func (r *medicalinfoResolver) Bloodtype(ctx context.Context, obj *models.Medicalinfo) (*string, error) {
	panic("not implemented")
}
func (r *medicalinfoResolver) Conditions(ctx context.Context, obj *models.Medicalinfo) ([]*models.Condition, error) {
	panic("not implemented")
}
func (r *medicalinfoResolver) Allergies(ctx context.Context, obj *models.Medicalinfo) ([]*models.Allergy, error) {
	panic("not implemented")
}
func (r *medicalinfoResolver) Vitals(ctx context.Context, obj *models.Medicalinfo) (*models.Vitals, error) {
	panic("not implemented")
}
func (r *medicalinfoResolver) Metadata(ctx context.Context, obj *models.Medicalinfo) (*models.Metadata, error) {
	panic("not implemented")
}

type metadataResolver struct{ *Resolver }

func (r *metadataResolver) Date(ctx context.Context, obj *models.Metadata) (*int, error) {
	panic("not implemented")
}
func (r *metadataResolver) Lastedit(ctx context.Context, obj *models.Metadata) (*int, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePatient(ctx context.Context, input NewPatient) (*models.Patient, error) {
	panic("not implemented")
}

type nextofkinResolver struct{ *Resolver }

func (r *nextofkinResolver) Name(ctx context.Context, obj *models.Nextofkin) (*string, error) {
	panic("not implemented")
}
func (r *nextofkinResolver) Relationship(ctx context.Context, obj *models.Nextofkin) (*string, error) {
	panic("not implemented")
}
func (r *nextofkinResolver) Phone(ctx context.Context, obj *models.Nextofkin) (*string, error) {
	panic("not implemented")
}
func (r *nextofkinResolver) Workplace(ctx context.Context, obj *models.Nextofkin) (*string, error) {
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
func (r *patientResolver) HospAdmin(ctx context.Context, obj *models.Patient) (*models.HospAdmin, error) {
	panic("not implemented")
}
func (r *patientResolver) ID(ctx context.Context, obj *models.Patient) (string, error) {
	panic("not implemented")
}
func (r *patientResolver) Parentid(ctx context.Context, obj *models.Patient) (*string, error) {
	panic("not implemented")
}
func (r *patientResolver) Nextofkin(ctx context.Context, obj *models.Patient) (*models.Nextofkin, error) {
	panic("not implemented")
}
func (r *patientResolver) Insurance(ctx context.Context, obj *models.Patient) ([]*models.Insurance, error) {
	panic("not implemented")
}
func (r *patientResolver) Medicalinfo(ctx context.Context, obj *models.Patient) (*models.Medicalinfo, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Patient(ctx context.Context, id string) ([]*models.Patient, error) {
	panic("not implemented")
}
func (r *queryResolver) HospAdmin(ctx context.Context, id string) ([]*models.HospAdmin, error) {
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

type vitalsResolver struct{ *Resolver }

func (r *vitalsResolver) Height(ctx context.Context, obj *models.Vitals) (*int, error) {
	panic("not implemented")
}
func (r *vitalsResolver) Weight(ctx context.Context, obj *models.Vitals) (*int, error) {
	panic("not implemented")
}
func (r *vitalsResolver) Pressure(ctx context.Context, obj *models.Vitals) (*int, error) {
	panic("not implemented")
}
func (r *vitalsResolver) Sugar(ctx context.Context, obj *models.Vitals) (*int, error) {
	panic("not implemented")
}
func (r *vitalsResolver) Heartrate(ctx context.Context, obj *models.Vitals) (*int, error) {
	panic("not implemented")
}
func (r *vitalsResolver) Respiration(ctx context.Context, obj *models.Vitals) (*int, error) {
	panic("not implemented")
}
func (r *vitalsResolver) Hb(ctx context.Context, obj *models.Vitals) (*string, error) {
	panic("not implemented")
}
