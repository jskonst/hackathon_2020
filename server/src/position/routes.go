package position

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"src/common"
	"src/database"
)

// InitializeRoutes
func InitializeRoutes(router *mux.Router, db *database.Database) {
	router.HandleFunc("/api/positions", getPositionHandler(db)).Methods("GET")
	router.HandleFunc("/api/positions", addPositionHandler(db)).Methods("POST")
}

// getPositionHandler ...
func getPositionHandler(db *database.Database) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		repository := NewPositionRepository(db)
		positions, err := repository.GetPositions()

		if err != nil {
			common.ErrorResponse(writer, err)
			return
		}

		common.JSONResponse(writer, positions, http.StatusOK)
	}
}

// addPositionHandler ...
func addPositionHandler(db *database.Database) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var position Position
		repository := NewPositionRepository(db)

		request.Body = http.MaxBytesReader(writer, request.Body, 1048576)
		err := json.NewDecoder(request.Body).Decode(&position)

		if err != nil {
			common.ErrorResponse(writer, err)
			return
		}

		if err = repository.AddPosition(position); err != nil {
			common.ErrorResponse(writer, err)
			return
		}
	}
}
