package main

import (
	"fmt"
	"log"
	"net/http"
)

var Server http.Handler

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

	routes.HandleFunc("GET /students/export", exportStudents)
	Server = routes
}

func root(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Curricular API version: 0.0.1")

	if err != nil {
		log.Fatalln("/ error", err)
	}
}

func getStudents(w http.ResponseWriter, r *http.Request) {

}

func searchStudents(w http.ResponseWriter, r *http.Request) {

}

func getTeachers(w http.ResponseWriter, r *http.Request) {

}

func searchTeachers(w http.ResponseWriter, r *http.Request) {

}

func getClasses(w http.ResponseWriter, r *http.Request) {

}

func getTeachersForClass(w http.ResponseWriter, r *http.Request) {

}

func getStudentsForTeacher(w http.ResponseWriter, r *http.Request) {

}

func exportStudents(w http.ResponseWriter, r *http.Request) {

}
