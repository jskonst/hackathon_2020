package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jskonst/hackathon_2020/server/common"
	"github.com/jskonst/hackathon_2020/server/config"
	"github.com/jskonst/hackathon_2020/server/database"
	"github.com/jskonst/hackathon_2020/server/position"
	"log"
	"net/http"
)

func getPointHandler(database *database.Database) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		repository := position.NewRepository(database)
		positions, err := repository.GetPositions()
		if err != nil {
			common.CreateErrorResponse(writer, common.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
			return
		}

		response, err := json.Marshal(positions)
		if err != nil {
			common.CreateErrorResponse(writer, common.APIError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(response)
	}
}

func main() {
	cfg, err := config.New("../.env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(cfg.ConnectionString)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/positions", getPointHandler(db))

	err = http.ListenAndServe(":3000", router)
	log.Fatal(err)
}
