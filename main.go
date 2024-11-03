package main

import (
	"fmt"
	"weather/app/storage"
)

func main() {
	// Run the SetupDB function from the storage package
	storage.SetupDB()
	storage.SetupCache()

	// Print a message to the console
	fmt.Println("Database setup complete")

}
