package model

import "time"

type Identity struct {
	Name  string
	Value string
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Country      string
	ZipCode      string
}

type ResidencyStatus int

const (
	Resident ResidencyStatus = iota
	NonResident
	Undermined
)

type Student struct {
	Identities []Identity
	FirstName  string
	LastName   string
	Addresses  []Address
	Birthdate  time.Time
	AgeInYears int
	Residency  ResidencyStatus
}
