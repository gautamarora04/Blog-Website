package database

import (
	"log"
	"os"

	"github.com/gautamarora04/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //declaring global variable for database connection

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //gorm is a library for golang for connecting to database
	if err != nil {
		panic("Could not connect to Database")
	}
	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
