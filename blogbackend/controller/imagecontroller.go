package controller

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

const reactPublicFolder = "D:/Blog Website-Frontend/reactlogin-master/public/"

// we will add random string in start of file name so that it will not create any overlap with images with same name

func randletter(n int) string {
	b := make([]rune, n)
	for i, _ := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func UploadImage(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["image"]
	filename := ""

	for _, file := range files {
		filename = randletter(5) + "-" + file.Filename
		if err := c.SaveFile(file, reactPublicFolder+"/upload/"+filename); err != nil {
			// return nil
			return c.JSON(fiber.Map{
				"url":     filename,
				"message": "error",
			})
		}
	}
	return c.JSON(fiber.Map{
		// "url": "localhost:3000/api/upload/" + filename,
		"url": filename,
	})
}
