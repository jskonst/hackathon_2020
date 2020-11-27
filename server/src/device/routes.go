package device

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"src/common"
	"src/database"
	"strconv"
)

// InitializeRoutes ...
func InitializeRoutes(router *mux.Router, db *database.Database) {
	router.HandleFunc("/devices", getDevicesHandler(db)).Methods("GET")
	router.HandleFunc("/devices/{id:[0-9]+}", getDeviceByIdHandler(db)).Methods("GET")
	router.HandleFunc("/devices", addDeviceHandler(db)).Methods("POST")
}

// getDevicesHandler ...
func getDevicesHandler(db *database.Database) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		repository := NewDeviceRepository(db)
		devices, err := repository.GetDevices()

		if err != nil {
			common.ErrorResponse(writer, err)
			return
		}

		common.JSONResponse(writer, devices, http.StatusOK)
	}
}

// getDeviceByIdHandler ...
func getDeviceByIdHandler(db *database.Database) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 32)

		repository := NewDeviceRepository(db)
		device, err := repository.GetDeviceById(int(id))

		if err != nil {
			common.ErrorResponse(writer, err)
			return
		}

		common.JSONResponse(writer, device, http.StatusOK)
	}
}

// addDeviceHandler ...
func addDeviceHandler(db *database.Database) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var device Device
		repository := NewDeviceRepository(db)

		request.Body = http.MaxBytesReader(writer, request.Body, 1048576)
		err := json.NewDecoder(request.Body).Decode(&device)

		if err != nil {
			common.ErrorResponse(writer, err)
			return
		}

		if err = repository.AddDevice(device); err != nil {
			common.ErrorResponse(writer, err)
			return
		}
	}
}
