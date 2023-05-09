package database

import (
	"log"
	"os"

	"github.com/hjuraev31/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn`t read file ")
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn`t connect to database")
	}
	DB = db
	db.AutoMigrate(
		&models.User{},
		&models.Cart{},
		&models.Products{},
	)
	return db
}
