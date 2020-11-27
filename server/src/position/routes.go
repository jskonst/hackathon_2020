package position

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jskonst/hackathon_2020/server/common"
	"github.com/jskonst/hackathon_2020/server/database"
	"net/http"
)

// InitializeRoutes
func InitializeRoutes(router *mux.Router, database *database.Database) {
	router.HandleFunc("/positions", getPositionHandler(database)).Methods("GET")
	router.HandleFunc("/positions", addPositionHandler(database)).Methods("POST")
}

// getPositionHandler ...
func getPositionHandler(database *database.Database) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		repository := NewRepository(database)
		positions, err := repository.GetPositions()

		if err != nil {
			common.APIErrorResponse(writer, common.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
			return
		}

		common.JSONResponse(writer, positions, http.StatusOK)
	}
}

// addPositionHandler ...
func addPositionHandler(database *database.Database) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var position Position
		repository := NewRepository(database)

		request.Body = http.MaxBytesReader(writer, request.Body, 1048576)
		err := json.NewDecoder(request.Body).Decode(&position)

		if err != nil {
			common.ErrorResponse(writer, err)
			return
		}

		if err = repository.AddPosition(&position); err != nil {
			common.ErrorResponse(writer, err)
			return
		}
	}
}
