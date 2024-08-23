package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes and returns a database connection
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./profiles.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
