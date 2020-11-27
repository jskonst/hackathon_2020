package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Database ...
type Database struct {
	*sqlx.DB
}

// New ...
func New(connectionString string) (*Database, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE EXTENSION IF NOT EXISTS postgis;")
	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}
