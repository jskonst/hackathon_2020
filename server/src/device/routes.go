package device

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"src/common"
	"src/database"
	"src/logger"
)

// InitializeRoutes ...
func InitializeRoutes(router *mux.Router, db *database.Database, logger *logger.Logger) {
	router.HandleFunc("/api/devices", getDevicesHandler(db, logger)).Methods("GET")
	router.HandleFunc("/api/devices/{imei}", getDeviceByIMEIHandler(db, logger)).Methods("GET")
	router.HandleFunc("/api/devices", addDeviceHandler(db, logger)).Methods("POST")
}

// getDevicesHandler ...
func getDevicesHandler(db *database.Database, logger *logger.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		repository := NewDeviceRepository(db)
		devices, err := repository.GetDevices()

		if err != nil {
			logger.Err(err)
			common.ErrorResponse(writer, err)
			return
		}

		common.JSONResponse(writer, devices, http.StatusOK)
	}
}

// getDeviceByIMEIHandler ...
func getDeviceByIMEIHandler(db *database.Database, logger *logger.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		imei := mux.Vars(request)["imei"]

		repository := NewDeviceRepository(db)
		device, err := repository.GetDeviceByIMEI(imei)

		if err != nil {
			logger.Err(err)
			common.ErrorResponse(writer, err)
			return
		}

		common.JSONResponse(writer, device, http.StatusOK)
	}
}

// addDeviceHandler ...
func addDeviceHandler(db *database.Database, logger *logger.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var device Device

		repository := NewDeviceRepository(db)
		request.Body = http.MaxBytesReader(writer, request.Body, 1048576)

		if err := json.NewDecoder(request.Body).Decode(&device); err != nil {
			logger.Err(err)
			common.ErrorResponse(writer, err)
			return
		}

		if err := repository.AddDevice(device); err != nil {
			logger.Err(err)
			common.ErrorResponse(writer, err)
			return
		}

		logger.Info().Interface("device", device).Msg("added new device")
	}
}
