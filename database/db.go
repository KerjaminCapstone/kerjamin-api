package database

import (
	"log"

	"github.com/KerjaminCapstone/kerjamin-backend-v1/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init(params ...string) *gorm.DB {
	conString := config.GetPostgresConnectionString()
	log.Print(conString)

	var err error
	db, err = gorm.Open(postgres.Open(conString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})

	if err != nil {
		log.Panic(err)
	}

	return db
}

func GetDBInstance() *gorm.DB {
	return db
}
