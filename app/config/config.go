package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppID      string
	AppKey     string
	MacAddress string
	DBName     string
	Interval   int
}

// Global variable to hold the loaded configuration
var cfg Config

// Load initializes the configuration by reading environment variables.
func Load() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Convert Interval from string to int
	intervalStr := os.Getenv("Interval")
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		return fmt.Errorf("invalid value for Interval: %s", intervalStr)
	}

	// Set the global cfg variable
	cfg = Config{
		AppID:      os.Getenv("AppID"),
		AppKey:     os.Getenv("AppKey"),
		MacAddress: os.Getenv("MacAddress"),
		DBName:     os.Getenv("DBName"),
		Interval:   interval,
	}

	// Optional: Check if required fields are set
	if cfg.AppID == "" || cfg.AppKey == "" || cfg.MacAddress == "" || cfg.DBName == "" {
		return fmt.Errorf("one or more required environment variables are not set")
	}

	return nil
}

// GetConfig returns the loaded configuration
func GetConfig() *Config {
	return &cfg
}
