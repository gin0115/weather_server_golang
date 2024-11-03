package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"weather/app/config"

	"github.com/boltdb/bolt"
)

type Storage struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"created_at"`
	ExpiresAt string `json:"expires_at"`
}

// SetupDB creates the database schema
func SetupCache() {
	// Get
	db := getDB()
	defer db.Close()

	// Create the tables if they don't exist
	createCacheBucket(db)
}

// Create the tables if they don't exist
func createCacheBucket(db *bolt.DB) {
	// Create the bucket if it doesn't exist
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("storage"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Get the bolt db instance.
func getDB() *bolt.DB {
	// Load and get the configuration.
	config.Load()
	config := config.GetConfig()

	// Check if DBName is set.
	if config.DBName == "" {
		log.Fatal("DB_NAME is not set")
	}

	// Create bolt database.
	db, err := bolt.Open("data/"+config.DBName+".cache", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// Add or update a key value pair in the cache
func SetCache(key string, value string, expires_at int) {
	db := getDB()
	defer db.Close()

	// If expires_at is no set, set it to 24 x 60 (Day in minutes)
	if expires_at == 0 {
		expires_at = 1440
	}

	// Create a new storage instance
	storage := Storage{
		Key:       key,
		Value:     value,
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(expires_at)).Format(time.RFC3339),
	}

	// Store the key value pair
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("storage"))
		encoded, err := json.Marshal(storage)
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), encoded)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

}

// Get a key value pair from the cache
func GetCache(key string) string {
	db := getDB()
	defer db.Close()

	var storage Storage

	// Get the key value pair
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("storage"))
		v := b.Get([]byte(key))
		if v == nil {
			return nil
		}
		err := json.Unmarshal(v, &storage)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Check if the key has expired based on its expires_at value
	if storage.ExpiresAt != "" {
		expiresAt, err := time.Parse(time.RFC3339, storage.ExpiresAt)
		if err != nil {
			log.Fatal(err)
		}
		if time.Now().After(expiresAt) {
			return ""
		}
	}

	return storage.Value
}
