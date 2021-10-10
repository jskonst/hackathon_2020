package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config ...
type Config struct {
	APIAddress string
	ListenAddress string
}

// New ...
func New(filenames ...string) (*Config, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		return nil, err
	}

	apiAddress, exists := os.LookupEnv("API_ADDRESS")
	if exists == false {
		return nil, fmt.Errorf("unable to connect to database: API_ADDRESS is empty")
	}

	listenAddress, exists := os.LookupEnv("LISTEN_ADDRESS")
	if exists == false {
		return nil, fmt.Errorf("unable to connect to database: LISTEN_ADDRESS is empty")
	}

	return &Config{APIAddress: apiAddress, ListenAddress: listenAddress}, nil
}
