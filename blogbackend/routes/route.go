package routes

import (
	"github.com/gautamarora04/controller"
	"github.com/gofiber/fiber/v2"
)

func Serve(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}
