package models

import "time"

type HospFile struct {
	id         string
	date       time.Time
	lastvisit  time.Time
	no         string
	idno       string
	visitcount int
}
