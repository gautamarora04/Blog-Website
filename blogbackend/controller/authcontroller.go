package controller

import (
	"log"
	"regexp"
	"strings"

	"github.com/gautamarora04/database"
	"github.com/gautamarora04/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// to check whether email is valid or not to save bunch of invalid records
func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

// to encrypt password before saving into database usihg golang bcrypt this is basically hashing, not encryting.
// For revison purpose, https://stackoverflow.com/questions/18084595/how-to-decrypt-hash-stored-by-bcrypt
func HashPassword(password string) []byte {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userdata models.User
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	// to check whether the password is greater than 8 char or not
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 characters",
		})
	}
	if !isEmailValid(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Email Address",
		})
	}
	// we need to check whether the email address already exists in database

	database.DB.Where("email = ?", strings.TrimSpace(data["email"].(string))).First(&userdata)
	if userdata.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email Address Already Exists",
		})

	}
	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Email:     data["email"].(string),
		Phone:     data["phone"].(string),
		Password:  HashPassword(data["password"].(string)),
	}
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account created Successfully",
	})

}
