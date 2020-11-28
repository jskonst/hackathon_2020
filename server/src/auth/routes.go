package auth

import (
	"context"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"src/common"
)

// InitializeRoutes ...
func InitializeRoutes(router *mux.Router, config *oauth2.Config) {
	router.HandleFunc("/api/auth", getAuthHandler(config)).Methods("GET")
	router.HandleFunc("/api/me", getMeHandler(config)).Methods("GET")
}

// getAuthHandler ...
func getAuthHandler(config *oauth2.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, config.AuthCodeURL("state", oauth2.AccessTypeOffline), http.StatusFound)
	}
}

// getMeHandler ...
func getMeHandler(config *oauth2.Config) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		code := request.FormValue("code")
		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			common.ErrorResponse(writer, err)
			return
		}

		client := config.Client(context.Background(), token)
		response, err := client.Get("https://www.googleapis.com/userinfo/v3/me")

		if err != nil {
			common.ErrorResponse(writer, err)
			return
		}

		defer response.Body.Close()

		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			common.ErrorResponse(writer, err)
			return
		}

		common.JSONResponse(writer, content, http.StatusOK)
	}
}
