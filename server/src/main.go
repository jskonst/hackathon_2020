package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"src/config"
	"src/database"
	"src/device"
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
	device.InitializeRoutes(router, db)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	err = http.ListenAndServe(":3000", c.Handler(router))
	log.Fatal(err)
}
