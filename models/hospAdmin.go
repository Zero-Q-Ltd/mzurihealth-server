package models

type AdminInvite struct {
	email      string
	name       string
	phone      string
	categoyid  string
	level      int
	inviterid  string
	hospitalid string
	metadata   Metadata
	id         string
}

type HospitalAdmin struct {
	id          string
	status      bool
	data        data
	config      adminconfig
	profiledata profiledata
	metadata    Metadata
}
type adminconfig struct {
	hospitalid   string
	categoryid   string
	level        int
	availability int // Whether on break , away or available
}

type profiledata struct {
	bio     string
	age     string
	address string
	phone   string
	status  bool // Whether olnine or offline
}
type data struct {
	uid         string
	email       string
	photoURL    string
	displayName string
}
