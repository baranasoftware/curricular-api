package model

import (
	"fmt"
	"time"
)

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

func (r ResidencyStatus) String() string {
	switch r {
	case Resident:
		return "Resident"
	case NonResident:
		return "NonResident"
	case Undermined:
		return "Undermined"
	default:
		return "Unknown"
	}
}

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

func (s Student) Print() {
	fmt.Printf("student: %+v\n", s)
}
