package position

import (
	"encoding/json"
	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
	"net/http"
	"src/common"
	"src/database"
	"src/logger"
)

// InitializeRoutes
func InitializeRoutes(router *mux.Router, db *database.Database, logger *logger.Logger, socket *socketio.Server) {
	router.HandleFunc("/api/positions", getPositionHandler(db, logger)).Methods("GET")
	router.HandleFunc("/api/positions/{imei}", getPositionByIMEI(db, logger)).Methods("GET")
	router.HandleFunc("/api/positions", addPositionHandler(db, logger, socket)).Methods("POST")
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

// getPositionByIMEI ...
func getPositionByIMEI(db *database.Database, logger *logger.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		imei := mux.Vars(request)["imei"]

		repository := NewPositionRepository(db)
		positions, err := repository.GetPositionsByIMEI(imei)

		if err != nil {
			logger.Err(err)
			common.ErrorResponse(writer, err)
			return
		}

		common.JSONResponse(writer, positions, http.StatusOK)
	}
}

// addPositionHandler ...
func addPositionHandler(db *database.Database, logger *logger.Logger, socket *socketio.Server) http.HandlerFunc {
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

		socket.BroadcastToRoom("", "map", "position:new", position)
		logger.Info().Interface("position", position).Msg("added new position")
	}
}
