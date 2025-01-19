package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Server http.Handler
var store DataStore

func init() {
	routes := http.NewServeMux()
	routes.HandleFunc("GET /", root)

	routes.HandleFunc("GET /students", getStudents)
	routes.HandleFunc("GET /students/search", searchStudents)

	routes.HandleFunc("GET /teachers", getTeachers)
	routes.HandleFunc("GET /teachers/search", searchTeachers)
	routes.HandleFunc("GET /teachers/{teacherId}/students", getStudentsForTeacher)

	routes.HandleFunc("GET /classes", getClasses)
	routes.HandleFunc("GET /classes/{classId}/teachers", getTeachersForClass)

	routes.HandleFunc("GET /courses", getCourses)
	routes.HandleFunc("GET /courses/{courseId}/students", getStudentsForCourse)

	routes.HandleFunc("GET /students/export", exportStudents)
	Server = routes

	numberOfRecords := 10
	store = NewDataStore(numberOfRecords)
}

func root(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Curricular API version: 0.0.1")

	if err != nil {
		log.Println("error: /", err)
	}
}

func getStudents(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(store.Students())
	if err != nil {
		log.Println("error: /students", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func searchStudents(w http.ResponseWriter, r *http.Request) {

}

func getTeachers(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(store.Teachers())
	if err != nil {
		log.Println("error: /teachers", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func searchTeachers(w http.ResponseWriter, r *http.Request) {

}

func getClasses(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(store.Classes())
	if err != nil {
		log.Println("error: /classes", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func getTeachersForClass(w http.ResponseWriter, r *http.Request) {

}

func getStudentsForTeacher(w http.ResponseWriter, r *http.Request) {

}

func getCourses(w http.ResponseWriter, r *http.Request) {

}

func getStudentsForCourse(w http.ResponseWriter, r *http.Request) {

}

func exportStudents(w http.ResponseWriter, r *http.Request) {

}
