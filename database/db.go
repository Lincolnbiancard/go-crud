package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/Lincolnbiancard/gin-api/models"
)
  
var (
	DB *gorm.DB
	err error
)

func ConnectToDB() {
	stringConnection := "host=172.19.0.2 user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConnection))
	if err != nil {
		log.Panic("Error connecting to database")
	}
	DB.AutoMigrate(&models.Student{})
}
