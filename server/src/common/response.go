package common

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIError ...
type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// CreateErrorResponse ...
func CreateErrorResponse(writer http.ResponseWriter, e APIError) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(e.StatusCode)

	response, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}
