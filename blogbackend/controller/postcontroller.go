package controller

import (
	"fmt"

	"github.com/gautamarora04/database"
	"github.com/gautamarora04/models"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var BlogPost models.Blog
	if err := c.BodyParser(&BlogPost); err != nil {
		fmt.Println("Error in body Parsing")
	}
	if err := database.DB.Create(&BlogPost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Error creating Payload",
		})
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Post Created Successfully",
	})
}
