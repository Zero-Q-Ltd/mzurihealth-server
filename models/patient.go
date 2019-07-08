package models

import "time"

type Patient struct {
	personalinfo personalinfo `json:"personalinfo"`
	fileinfo     HospFile     `json:"fileinfo"`
	id           string       `json:"id"`
	/**
	 * Optional parent id int for minors
	 */
	parentid string `json:"parentid"`

	nextofkin nextofkin `json:"nextofkin"`
	/**
	 * A patient can have several insurances at the same time
	 */
	insurance   []Insurance `json:"isnurance"`
	medicalinfo medicalinfo `json:"medicalinfo"`
	/**
	 * used in queries so that you can optionally disable some patients
	 */
	status       bool           `json:"status"`
	exrainfo     string         `json:"extarinfo"`
	customfuelds []Customfields `json:"customfields"`
	primaryhosp  string         `json:"primaryhosp"`
	metadata     Metadata       `json:"metadata"`
}

type Insurance struct {
	id          string `json:"gender"`
	insuranceno string `json:"gender"`
}
type Nextofkin struct {
	name         string `json:"gender"`
	relationship string `json:"gender"`
	phone        int    `json:"gender"`
	workplace    string `json:"gender"`
}

type Medicalinfo struct {
	bloodtype  string      `json:"bloodtype"`
	conditions []Condition `json:"condition"`
	allergies  []Allegy    `json:"allergies"`
	vitals     vitals      `json:"vitals"`
	metadata   *Metadata   `json:"metadata"`
}

type personalinfo struct {
	address    string    `json:"address,omitempty"`
	photoURL   string    `json:"photourl"`
	name       string    `json:"name"`
	gender     int       `json:"gender"`
	occupation string    `json:"occupation"`
	workplace  string    `json:"workplace"`
	phone      int       `json:"phone"`
	email      string    `json:"email"`
	idno       string    `json:"idno"`
	dob        time.Time `json:"dob"`
}
type Vitals struct {
	height      int    `json:"height"`
	weight      int    `json:"weight"`
	pressure    int    `json:"pressure"`
	sugar       int    `json:"sugar"`
	heartrate   int    `json:"heartrate"`
	respiration int    `json:"respiration"`
	hb          string `json:"hb"`
}
