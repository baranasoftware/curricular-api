package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/server"
	"log"
	"net/http"
)

type Configuration struct {
	dataStore    DataStore
	oauth2Server *server.Server
	totalRecords int
}

func NewConfiguration(n int) Configuration {
	ds := NewDataStore(n)
	os, err := NewOAuth2Manager()
	if err != nil {
		log.Fatalf("cloud not create OAuth2 Manager: %s", err)
	}

	return Configuration{ds, os, n}
}

func Handlers(c Configuration) http.Handler {
	// for AWS APIGateway routes are mapped using https://github.com/awslabs/aws-lambda-go-api-proxy/
	routes := http.NewServeMux()
	routes.HandleFunc("GET /", c.root)

	routes.HandleFunc("GET /students", c.getStudents)
	routes.HandleFunc("GET /students/search", c.searchStudents)

	routes.HandleFunc("GET /teachers", c.getTeachers)
	routes.HandleFunc("GET /teachers/search", c.searchTeachers)
	routes.HandleFunc("GET /teachers/{teacherId}/students", c.getStudentsForTeacher)

	routes.HandleFunc("GET /classes", c.getClasses)
	routes.HandleFunc("GET /classes/{classId}/teachers", c.getTeachersForClass)

	routes.HandleFunc("GET /courses", c.getCourses)
	routes.HandleFunc("GET /courses/{courseId}/students", c.getStudentsForCourse)

	routes.HandleFunc("GET /students/export", c.exportStudents)

	// only for demonstration purposes, in actual setup will be implemented by
	// the APIGateway -- Apigee
	routes.HandleFunc("GET /authorize", c.authorize)
	routes.HandleFunc("GET /oauth/token", c.token)

	return routes
}

func (c Configuration) root(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Curricular API version: 0.0.1")

	if err != nil {
		log.Println("error: /", err)
	}
}

func (c Configuration) getStudents(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(c.dataStore.Students())
	if err != nil {
		log.Println("error: /students", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (c Configuration) searchStudents(w http.ResponseWriter, r *http.Request) {

}

func (c Configuration) getTeachers(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(c.dataStore.Teachers())
	if err != nil {
		log.Println("error: /teachers", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (c Configuration) searchTeachers(w http.ResponseWriter, r *http.Request) {

}

func (c Configuration) getClasses(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(c.dataStore.Classes())
	if err != nil {
		log.Println("error: /classes", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (c Configuration) getTeachersForClass(w http.ResponseWriter, r *http.Request) {

}

func (c Configuration) getStudentsForTeacher(w http.ResponseWriter, r *http.Request) {

}

func (c Configuration) getCourses(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(c.dataStore.Courses())
	if err != nil {
		log.Println("error: /courses", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (c Configuration) getStudentsForCourse(w http.ResponseWriter, r *http.Request) {

}

func (c Configuration) exportStudents(w http.ResponseWriter, r *http.Request) {

}
