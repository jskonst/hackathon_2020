package auth

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"net/http"
)

// InitializeRoutes ...
func InitializeRoutes(router *mux.Router, config *oauth2.Config) {
	router.HandleFunc("/api/auth", getAuthHandler(config)).Methods("GET")
	router.HandleFunc("/api/me", getMeHandler()).Methods("GET")
}

// getAuthHandler ...
func getAuthHandler(config *oauth2.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, config.AuthCodeURL("state", oauth2.AccessTypeOffline), http.StatusFound)
	}
}

// getMeHandler ...
func getMeHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.GetBody())
	}
}
