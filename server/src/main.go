package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"src/config"
	"src/database"
	"src/position"
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
