package main

import (
	socketio "github.com/googollee/go-socket.io"
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

	socket, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	socket.OnConnect("/", func(conn socketio.Conn) error {
		log.Println("NEW CONNECTION")

		return nil
	})

	position.InitializeRoutes(router, db, logg)
	device.InitializeRoutes(router, db, logg)
	auth.InitializeRoutes(router, ocfg)

	//router.Handle("/socket.io", socket)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "GET", "DELETE"},
		AllowCredentials: true,
	})

	err = http.ListenAndServe(":3000", c.Handler(router))
	log.Fatal(err)
}
