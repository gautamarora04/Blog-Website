package main

import (
	"log"
	"os"

	"github.com/gautamarora04/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	app.Listen(":" + port)

}
