package database

import (
	"log"
	"os"

	"github.com/itaypanda/fiber-auth-example/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	connectionString := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("An error ocurred while connecting to the database: %v", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("An error ocurred while migrating the database: %v", err)
	}
	log.Print("Database connected!")
}
