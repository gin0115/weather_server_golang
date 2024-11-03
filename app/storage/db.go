package storage

import (
	"database/sql"
	"log"
	"weather/app/config"

	_ "github.com/mattn/go-sqlite3"
)

// Gets the DB instance.
func GetDB() *sql.DB {
	// Load and get the configuration.
	config.Load()
	config := config.GetConfig()

	// Check if DBName is set.
	if config.DBName == "" {
		log.Fatal("DB_NAME is not set")
	}

	// Create sqlite database.
	db, err := sql.Open("sqlite3", "data/"+config.DBName+".db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
