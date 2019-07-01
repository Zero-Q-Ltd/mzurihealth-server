package models

import "time"

type Customfields struct {
	id    int
	value struct{}
	name  string
}

type Metadata struct {
	date     time.Time
	lastedit time.Time
}
