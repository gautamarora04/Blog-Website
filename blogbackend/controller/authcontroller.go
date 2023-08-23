package controller

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gautamarora04/database"
	"github.com/gautamarora04/models"
	"github.com/gautamarora04/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// to check whether email is valid or not to save bunch of invalid records
func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
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
	}
	user.HashPassword(data["password"].(string))
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

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email Id is not registed, Kindly create account",
		})
	}
	err := user.CheckPassword(data["password"])

	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Wrong Password",
		})
	}
	token, err := util.GenerateJWT(strconv.Itoa(int(user.ID)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	c.Status(200)

	return c.JSON(fiber.Map{
		"message": "Login Successful",
		"user":    user,
	})
}

type Claims struct {
	jwt.StandardClaims
}
