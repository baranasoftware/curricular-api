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

	Server = routes
}

func root(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Curricular API version: 0.0.1")

	if err != nil {
		log.Fatalln("/ error", err)
	}
}
