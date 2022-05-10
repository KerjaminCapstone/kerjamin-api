package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type MySQLConnString struct {
	dbUsername string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
}

func getMySQLCred() MySQLConnString {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error to load .env file!")
	}

	return MySQLConnString{
		dbUsername: os.Getenv("DB_USERNAME"),
		dbPassword: os.Getenv("DB_PASSWORD"),
		dbName:     os.Getenv("DB_DATABASE"),
		dbHost:     os.Getenv("DB_HOST"),
		dbPort:     os.Getenv("DB_PORT"),
	}
}

func GetMySQLConnString() string {
	credential := getMySQLCred()
	database := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		credential.dbUsername,
		credential.dbPassword,
		credential.dbHost,
		credential.dbPort,
		credential.dbName,
	)

	return database
}
