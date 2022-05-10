package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PostgresCredential struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func GetPostgresCredential() PostgresCredential {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error to load .env file")
	}

	return PostgresCredential{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_DATABASE"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}

func GetPostgresConnectionString() string {
	credential := GetPostgresCredential()
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		credential.DBHost,
		credential.DBPort,
		credential.DBUsername,
		credential.DBName,
		credential.DBPassword,
	)
	return dataBase
}
