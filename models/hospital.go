package models

type Hospital struct {
	location       struct{}
	name           string
	userid         string
	_id            string
	description    string
	status         bool
	contacPerson   contactperson
	contactDetails contactDetails
	logourl        string
	patientCount   int
	invoiceCount   int
	metadata       Metadata
	paymentMethods []CustomPaymentMethod
}

type contactperson struct {
	name     string
	phone    string
	email    string
	position string
	address  string
}

type contactDetails struct {
	phone   string
	email   string
	address string
}
