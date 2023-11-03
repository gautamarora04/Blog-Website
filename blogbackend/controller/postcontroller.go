package controller

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gautamarora04/database"
	"github.com/gautamarora04/models"
	"github.com/gautamarora04/util"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var BlogPost models.Blog
	if err := c.BodyParser(&BlogPost); err != nil {
		fmt.Println("Error in body Parsing")
	}
	currentTime := time.Now()
	currentTimeStr := currentTime.Format("2006-01-02 15:04:05")
	BlogPost.CreatedAt = currentTimeStr
	if err := database.DB.Create(&BlogPost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Error creating Payload",
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Post Created Successfully",
		"time":    BlogPost,
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 6
	offset := (page - 1) * limit //Offset specify the number of records to skip before starting to return the records
	var total int64              // this variable is used to return the total number of pages to frontend
	var getblog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.Blog{}).Count(&total)

	last := math.Ceil(float64(int(total) / limit))
	if int(total)%limit > 0 {
		last = last + 1
	}

	return c.JSON(fiber.Map{
		"data": getblog,
		"meta": fiber.Map{
			"page":      page,
			"total":     total,
			"last_page": last,
		},
	})
}

func Detailpost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		ID: uint(id),
	}
	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}
	database.DB.Model(&blog).Updates(blog).Where("id=?", id).Preload("User").First(&blog)
	return c.JSON(fiber.Map{
		"message": blog,
	})
}

func MyPost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJWT(cookie)
	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)
	return c.JSON(blog)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blog models.Blog

	database.DB.Where("id=?", id).Preload("User").First(&blog)

	if blog.User.ID == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Record Not found!!",
		})
	}
	database.DB.Delete(&blog)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Post Deleted Successfully",
	})

}
