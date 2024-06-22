package internal

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
