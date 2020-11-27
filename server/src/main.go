package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"src/auth"
	"src/config"
	"src/database"
	"src/device"
	"src/logger"
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

	ocfg := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  "http://localhost:3000/api/me",
		Scopes:       []string{"profile"},
		Endpoint:     google.Endpoint,
	}

	logg := logger.NewLogger()
	router := mux.NewRouter()

	position.InitializeRoutes(router, db, logg)
	device.InitializeRoutes(router, db, logg)
	auth.InitializeRoutes(router, ocfg)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	err = http.ListenAndServe(":3000", c.Handler(router))
	log.Fatal(err)
}
