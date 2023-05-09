package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hjuraev31/database"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Couldn`t read .env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	log.Fatal(app.Listen(":" + port))
}
