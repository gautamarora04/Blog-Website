package main

import (
	"log"
	"os"

	"github.com/gautamarora04/database"
	"github.com/gautamarora04/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	port := os.Getenv("PORT")
	// CORS middleware configuration
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Replace with your frontend URL(s)
		AllowMethods:     "GET,POST,PUT,DELETE,USE,STATIC",
		AllowHeaders:     "Authorization,Content-Type",
		AllowCredentials: true, // Allow credentials (cookies, HTTP authentication) to be sent with the request
	}))
	routes.Serve(app)
	app.Listen(":" + port)

}
