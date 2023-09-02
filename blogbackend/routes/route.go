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
	app.Post("/api/logout", controller.Logout)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
	app.Get("/api/allpost/:id", controller.Detailpost)
	app.Put("/api/update/:id", controller.UpdatePost)
	app.Get("/api/mypost", controller.MyPost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
	app.Post("/api/upload-image", controller.UploadImage)
	app.Static("/api/upload/", "./upload/")
}
