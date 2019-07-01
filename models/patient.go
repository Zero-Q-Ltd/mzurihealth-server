package models

import "time"

type Patient struct {
	Personalinfo personalinfo `json:"personalinfo"`
	Fileinfo     HospFile     `json:"fileinfo"`
	Id           string       `json:"id"`
	/**
	 * Optional parent id int for minors
	 */
	Parentid string `json:"parentid"`

	Nextofkin nextofkin `json:"nextofkin"`
	/**
	 * A patient can have several insurances at the same time
	 */
	Insurance   []Insurance `json:"isnurance"`
	Medicalinfo medicalinfo `json:"medicalinfo"`
	/**
	 * used in queries so that you can optionally disable some patients
	 */
	Status       bool           `json:"status"`
	Exrainfo     string         `json:"extarinfo"`
	Customfuelds []Customfields `json:"customfields"`
	Primaryhosp  string         `json:"primaryhosp"`
	Metadata     Metadata       `json:"metadata"`
}

type Insurance struct {
	id          string `json:"gender"`
	insuranceno string `json:"gender"`
}
type nextofkin struct {
	Name         string `json:"gender"`
	Relationship string `json:"gender"`
	Phone        int    `json:"gender"`
	Workplace    string `json:"gender"`
}

type medicalinfo struct {
	Bloodtype  string      `json:"bloodtype"`
	Conditions []Condition `json:"condition"`
	Allergies  []Allegy    `json:"allergies"`
	Vitals     vitals      `json:"vitals"`
	Metadata   *Metadata   `json:"metadata"`
}

type personalinfo struct {
	Address    string    `json:"address,omitempty"`
	PhotoURL   string    `json:"photourl"`
	Name       string    `json:"name"`
	Gender     int       `json:"gender"`
	Occupation string    `json:"occupation"`
	Workplace  string    `json:"workplace"`
	Phone      int       `json:"phone"`
	Email      string    `json:"email"`
	Idno       string    `json:"idno"`
	Dob        time.Time `json:"dob"`
}
type vitals struct {
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
	Pressure    int    `json:"pressure"`
	Sugar       int    `json:"sugar"`
	Heartrate   int    `json:"heartrate"`
	Respiration int    `json:"respiration"`
	Hb          string `json:"hb"`
}
