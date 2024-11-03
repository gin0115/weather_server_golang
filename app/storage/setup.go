package storage

import (
	"database/sql"
	"log"
)

// SetupDB creates the database schema
func SetupDB() {
	// Get
	db := GetDB()
	defer db.Close()

	// Create the tables if they don't exist
	createTables(db)
}

// Create the tables if they don't exist
func createTables(db *sql.DB) {
	// Temperature table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS temperature (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		temperature FLOAT,
		feels_like FLOAT,
		dew_point FLOAT,
		apparent_temperature FLOAT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Solar table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS solar (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		radiation FLOAT,
		uv_index FLOAT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Rainfall table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS rainfall (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		rain_rate FLOAT,
		daily_rain FLOAT,
		event_rain FLOAT,
		hourly_rain FLOAT,
		weekly_rain FLOAT,
		monthly_rain FLOAT,
		yearly_rain FLOAT
	)`)
	if err != nil {
		log.Fatal(err)
	}
	// Wind table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS wind (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		wind_speed FLOAT,
		wind_gust FLOAT,
		wind_direction FLOAT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// atmospheric pressure table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS atmospheric (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		pressure_relative FLOAT,
		pressure_absolute FLOAT,
		humidity FLOAT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the weather records table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS weather_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		temperature FLOAT, 
		solar FLOAT, 
		rainfall FLOAT, 
		wind FLOAT, 
		atmospheric FLOAT 
	)`)
	if err != nil {
		log.Fatal(err)
	}

}
