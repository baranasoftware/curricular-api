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
	var teachers []model.Teacher
	var classes []model.Class

	idNames := map[int]string{0: "campusId", 1: "emplId", 2: "libraryId"}

	residency := []model.ResidencyStatus{model.Resident, model.NonResident, model.Undermined}
	type course struct {
		name   string
		credit float64
	}
	allClasses := map[int]course{
		0: {
			name:   "PROGRAMMING I",
			credit: 3.00,
		},
		1: {
			name:   "DATA SCIENCE PROGRAMMING I",
			credit: 4.00,
		},
		2: {
			name:   "INTRODUCTION TO DISCRETE MATHEMATICS",
			credit: 3.00,
		},
		3: {
			name:   "NUMERICAL LINEAR ALGEBRA",
			credit: 3.00,
		},
		4: {
			name:   "SOFTWARE ENGINEERING",
			credit: 3.00,
		},
		5: {
			name:   "CONSTRUCTION OF COMPILERS",
			credit: 3.00,
		},
		6: {
			name:   "PRINCIPLES OF PROGRAMMING LANGUAGES",
			credit: 3.00,
		},
	}

	for i := 0; i < size; i++ {

		var studentIdentities []model.Identity
		for _, idName := range idNames {
			guid := xid.New()
			studentIdentities = append(studentIdentities, model.Identity{
				Name:  idName,
				Value: guid.String(),
			})
		}

		var studentAddresses []model.Address
		studentAddress1 := gofakeit.Address()
		studentAddress2 := gofakeit.Address()
		studentAddresses = append(studentAddresses,
			model.Address{
				AddressLine1: studentAddress1.Street,
				City:         studentAddress1.City,
				State:        studentAddress1.State,
				Country:      studentAddress1.Country,
				ZipCode:      studentAddress1.Zip,
			}, model.Address{
				AddressLine1: studentAddress2.Street,
				City:         studentAddress2.City,
				State:        studentAddress2.State,
				Country:      studentAddress2.Country,
				ZipCode:      studentAddress2.Zip,
			})

		dob := gofakeit.Date()
		students = append(students, model.Student{
			Identities: studentIdentities,
			Addresses:  studentAddresses,
			FirstName:  gofakeit.FirstName(),
			LastName:   gofakeit.LastName(),
			Birthdate:  dob,
			AgeInYears: time.Now().Year() - dob.Year(),
			Residency:  residency[rand.Intn(len(residency))],
		})

		var teacherIdentities []model.Identity
		for _, idName := range idNames {
			guid := xid.New()
			teacherIdentities = append(teacherIdentities, model.Identity{
				Name:  idName,
				Value: guid.String(),
			})
		}

		teachers = append(teachers, model.Teacher{
			Identities: teacherIdentities,
			FirstName:  gofakeit.FirstName(),
			LastName:   gofakeit.LastName(),
		})

		cls := allClasses[rand.Intn(len(allClasses))]
		classAddress := gofakeit.Address()
		classes = append(classes, model.Class{
			Id:     xid.New().String(),
			Name:   cls.name,
			Credit: model.NewCredit(cls.credit),
			Location: model.Address{
				AddressLine1: classAddress.Street,
				City:         classAddress.City,
				State:        classAddress.State,
				Country:      classAddress.Country,
				ZipCode:      classAddress.Zip,
			},
			Time: gofakeit.FutureDate(),
		})
	}
	ds.students = students
	ds.teachers = teachers
	ds.classes = classes

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
