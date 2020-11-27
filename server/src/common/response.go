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

// JSONResponse ...
func JSONResponse(writer http.ResponseWriter, response interface{}, statusCode int) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(statusCode)

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		log.Fatal(err)
	}
}

// CreateErrorResponse ...
func APIErrorResponse(writer http.ResponseWriter, e APIError) {
	JSONResponse(writer, e, e.StatusCode)
}

// ErrorResponse ...
func ErrorResponse(writer http.ResponseWriter, err error) {
	APIErrorResponse(writer, APIError{
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	})
}
