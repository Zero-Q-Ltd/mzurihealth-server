package models

// type MedicalConditionsModel string =
//     "Alzheimer\"s"
//     | "Arthritis"
//     | "Asthma"
//     | "Blood Pressure"
//     | "Cancer"
//     | "Cholesterol"
//     | "Chronic Pain"
//     | "Cold & Flu"
//     | "Depression"
//     | "Diabetes"
//     | "Dictionary"
//     | "Digestion"
//     | "Eyesight"
//     | "Health & Living"
//     | "Healthy Kids"
//     | "Hearing & Ear"
//     | "Heart"
//     | "HIV/AIDS"
//     | "Infectious Disease"
//     | "Lung Conditions"
//     | "Medications"
//     | "Menopause"
//     | "Men\"s Health"
//     | "Mental"
//     | "Health"
//     | "Migraine"
//     | "Neurology"
//     | "Oral"
//     | "Health"
//     | "Pregnancy"
//     | "Senior"
//     | "Health"
//     | "Sexual"
//     | "Health"
//     | "Skin"
//     | "Sleep"
//     | "Thyroid"
//     | "Travel"
//     | "Health"
// 	| "Women\"s Health" ;

type Condition struct {
	conditiontype string
	detail        string
	metadata      Metadata
}

type Allergy struct {
	allergytype string
	detail      string
	metadata    Metadata
}

type RawProcedure struct {
	name      string
	_id       string
	pricing   pricing
	category  ProcedureCategory
	numericid int32
}
type ProcedureCategory struct {
	_id           string
	code          string
	subcategoryid string
}
type pricing struct {
	min int32
	max int32
}
