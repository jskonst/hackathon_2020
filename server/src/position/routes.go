package position

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"src/common"
	"src/database"
	"src/logger"
)

// InitializeRoutes
func InitializeRoutes(router *mux.Router, db *database.Database, logger *logger.Logger) {
	router.HandleFunc("/api/positions", getPositionHandler(db, logger)).Methods("GET")
	router.HandleFunc("/api/positions", addPositionHandler(db, logger)).Methods("POST")
}

// getPositionHandler ...
func getPositionHandler(db *database.Database, logger *logger.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		repository := NewPositionRepository(db)
		positions, err := repository.GetPositions()

		if err != nil {
			logger.Err(err)
			common.ErrorResponse(writer, err)
			return
		}

		common.JSONResponse(writer, positions, http.StatusOK)
	}
}

// addPositionHandler ...
func addPositionHandler(db *database.Database, logger *logger.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var position AddPositionRequestModel
		repository := NewPositionRepository(db)

		request.Body = http.MaxBytesReader(writer, request.Body, 1048576)
		err := json.NewDecoder(request.Body).Decode(&position)

		if err != nil {
			logger.Err(err)
			common.ErrorResponse(writer, err)
			return
		}

		if err = repository.AddPositionByIMEI(position); err != nil {
			logger.Err(err)
			common.ErrorResponse(writer, err)
			return
		}

		logger.Info().Interface("position", position).Msg("added new position")
	}
}
