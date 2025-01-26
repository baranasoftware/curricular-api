package api

import (
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"log"
	"net/http"
)

// AuthNZ functionality is here only for demonstration purposes.
// In the actual setup, this will be provided by the API Gateway product (e.g. Apigee)

func NewOAuth2Manager() *server.Server {
	manager := manage.NewDefaultManager()

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	return srv
}

func authorize(w http.ResponseWriter, r *http.Request) {
	err := oauth2Server.HandleAuthorizeRequest(w, r)
	if err != nil {
		log.Println("error: /authorize", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func token(w http.ResponseWriter, r *http.Request) {
	err := oauth2Server.HandleTokenRequest(w, r)
	if err != nil {
		log.Println("error: /token", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
