package main

import (
	"github.com/gorilla/mux"
	"github.com/jskonst/hackathon_2020/server/config"
	"github.com/jskonst/hackathon_2020/server/database"
	"github.com/jskonst/hackathon_2020/server/position"
	"log"
	"net/http"
)

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
	position.InitializeRoutes(router, db)

	err = http.ListenAndServe(":3000", router)
	log.Fatal(err)
}
