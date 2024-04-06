package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var StoragePath string

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	StoragePath = os.Getenv("STORAGE")
}
