package model

import "time"

type Teacher struct {
	Identities []Identity `json:"identities"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Addresses  []Address  `json:"addresses"`
	Birthdate  time.Time  `json:"birthDate"`
	AgeInYears int        `json:"ageInYears"`
}
