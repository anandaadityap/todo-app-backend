package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo-app/controller"
)

func RouteIndex(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", controller.GetAll)
	api.Post("/", controller.Create)
	api.Get("/:id", controller.GetById)
	api.Put("/:id", controller.Update)
	api.Delete("/:id", controller.Delete)
}
