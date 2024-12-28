package model

import "time"

type Teacher struct {
	Identities []Identity
	FirstName  string
	LastName   string
	Addresses  []Address
	Birthdate  time.Time
	AgeInYears int
}
