package database

import (
	"log"

	"github.com/xdouglas90/go_studies/gin-rest-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	err    error
)

func Connect() {
	dsn := "host=localhost user=root password=root dbname=root port=5434 sslmode=disable TimeZone=America/Sao_Paulo"
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	DBConn.AutoMigrate(&models.Student{})
}
