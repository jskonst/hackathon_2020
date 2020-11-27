package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config ...
type Config struct {
	ConnectionString string
	ClientID         string
	ClientSecret     string
}

// New ...
func New(filenames ...string) (*Config, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		return nil, err
	}

	connectionString, exists := os.LookupEnv("CONNECTION_STRING")
	if exists == false {
		return nil, fmt.Errorf("unable to connect to database: CONNECTION_STRING is empty")
	}

	clientId, exists := os.LookupEnv("CLIENT_ID")
	if exists == false {
		return nil, fmt.Errorf("")
	}

	clientSecret, exists := os.LookupEnv("CLIENT_SECRET")
	if exists == false {
		return nil, fmt.Errorf("")
	}

	return &Config{
		ConnectionString: connectionString,
		ClientID:         clientId,
		ClientSecret:     clientSecret,
	}, nil
}
