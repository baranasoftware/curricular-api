package api

import (
	"curricular-api/model"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/rs/xid"
	"math/rand"
	"time"
)

type DataStore struct {
	students []model.Student
	teachers []model.Teacher
	classes  []model.Class
	courses  []model.Course
}

func NewDataStore(size int) DataStore {

	ds := DataStore{}

	var students []model.Student
	idNames := map[int]string{0: "campusId", 1: "emplId", 2: "libraryId"}
	residency := []model.ResidencyStatus{model.Resident, model.NonResident, model.Undermined}
	for i := 0; i < size; i++ {

		var identities []model.Identity
		for _, idName := range idNames {
			guid := xid.New()
			identities = append(identities, model.Identity{
				Name:  idName,
				Value: guid.String(),
			})
		}

		var addresses []model.Address
		address1 := gofakeit.Address()
		address2 := gofakeit.Address()
		addresses = append(addresses,
			model.Address{
				AddressLine1: address1.Street,
				City:         address1.City,
				State:        address1.State,
				Country:      address1.Country,
				ZipCode:      address1.Zip,
			}, model.Address{
				AddressLine1: address2.Street,
				City:         address2.City,
				State:        address2.State,
				Country:      address2.Country,
				ZipCode:      address2.Zip,
			})

		dob := gofakeit.Date()
		students = append(students, model.Student{
			Identities: identities,
			Addresses:  addresses,
			FirstName:  gofakeit.FirstName(),
			LastName:   gofakeit.LastName(),
			Birthdate:  dob,
			AgeInYears: time.Now().Year() - dob.Year(),
			Residency:  residency[rand.Intn(len(residency))],
		})
	}

	ds.students = students

	return ds
}

func (d DataStore) Students() []model.Student {
	return d.students
}

func (d DataStore) Teachers() []model.Teacher {
	return d.teachers
}

func (d DataStore) Classes() []model.Class {
	return d.classes
}

func (d DataStore) Courses() []model.Course {
	return d.courses
}
