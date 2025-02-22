package api

import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"log"
	"net/http"
)

// AuthNZ functionality is here only for demonstration purposes.
// In the actual setup, this will be provided by the API Gateway product (e.g. Apigee or AWS APIGateway).
// So that we only have to focus on implementing logic specific to our use cases (implementing REST resources).

func NewOAuth2Manager() (*server.Server, error) {
	manager := manage.NewDefaultManager()

	// client memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore()) // in-memory token store

	// demonstration purposes
	clientId := "000000"
	secret := "999999"
	domain := "http://localhost"

	clientStore := store.NewClientStore()
	err := clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: secret,
		Domain: domain,
	})
	if err != nil {
		return nil, err
	}

	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.UserAuthorizationHandler = func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		return clientId, nil
	}

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("internal error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("response error:", re.Error.Error())
	})

	return srv, nil
}

func (c Configuration) register(w http.ResponseWriter, r *http.Request) {

}

func (c Configuration) authorize(w http.ResponseWriter, r *http.Request) {
	err := c.oauth2Server.HandleAuthorizeRequest(w, r)
	if err != nil {
		log.Println("error: /authorize", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (c Configuration) token(w http.ResponseWriter, r *http.Request) {
	err := c.oauth2Server.HandleTokenRequest(w, r)
	if err != nil {
		log.Println("error: /oauth/token", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
