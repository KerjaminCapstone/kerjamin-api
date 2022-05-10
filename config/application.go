package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetApplicationPort() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error to load .env file")
	}

	return os.Getenv("APP_PORT")
}
