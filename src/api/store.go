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

		dob := gofakeit.Date()
		students = append(students, model.Student{
			Identities: identities,
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
