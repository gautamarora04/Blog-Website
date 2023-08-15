package routes

import (
	"github.com/gautamarora04/controller"
	"github.com/gautamarora04/middleware"
	"github.com/gofiber/fiber/v2"
)

func Serve(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Use(middleware.IsAuthenticate)
	app.Post("/api/post", controller.CreatePost)

}
