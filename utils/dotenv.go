package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GodotEnv loads environment variables from the .env file and returns the value of the specified key.
func GodotEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return os.Getenv(key)
}
