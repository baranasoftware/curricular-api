package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/server"
	"log"
	"net/http"
)

var Server http.Handler
var dataStore DataStore
var oauth2Server *server.Server
var LocalSetup bool

func init() {
	// for AWS APIGateway routes are mapped using https://github.com/awslabs/aws-lambda-go-api-proxy/
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

	// only for demonstration purposes, in actual setup will be implemented by
	// the APIGateway -- Apigee
	routes.HandleFunc("GET /authorize", authorize)
	routes.HandleFunc("GET /oauth/token", token)

	Server = routes

	numberOfRecords := 20
	dataStore = NewDataStore(numberOfRecords)
	oauth2Server, _ = NewOAuth2Manager()
}

func root(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Curricular API version: 0.0.1")

	if err != nil {
		log.Println("error: /", err)
	}
}

func getStudents(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(dataStore.Students())
	if err != nil {
		log.Println("error: /students", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func searchStudents(w http.ResponseWriter, r *http.Request) {

}

func getTeachers(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(dataStore.Teachers())
	if err != nil {
		log.Println("error: /teachers", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func searchTeachers(w http.ResponseWriter, r *http.Request) {

}

func getClasses(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(dataStore.Classes())
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
	err := json.NewEncoder(w).Encode(dataStore.Courses())
	if err != nil {
		log.Println("error: /courses", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func getStudentsForCourse(w http.ResponseWriter, r *http.Request) {

}

func exportStudents(w http.ResponseWriter, r *http.Request) {

}
