package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PostgresCredential struct {
	DBUsername             string
	DBPassword             string
	DBName                 string
	InstanceConnectionName string
}

func GetPostgresCredential() PostgresCredential {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error to load .env file")
	}

	return PostgresCredential{
		DBUsername:             os.Getenv("DB_USERNAME"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBName:                 os.Getenv("DB_DATABASE"),
		InstanceConnectionName: os.Getenv("INSTANCE_CONNECTION_NAME"),
	}
}

func GetPostgresConnectionString() string {
	credential := GetPostgresCredential()
	dataBase := fmt.Sprintf("user=%s password=%s database=%s host=/cloudsql/%s",
		credential.DBUsername,
		credential.DBPassword,
		credential.DBName,
		credential.InstanceConnectionName,
	)
	return dataBase
}
