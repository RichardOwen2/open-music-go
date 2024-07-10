package routes

import (
	"openmusic-api/controller"

	"github.com/gofiber/fiber/v2"
)

func AlbumRoutes(app *fiber.App, c controller.AlbumController) {
	group := app.Group("/albums")

	group.Post("/", c.Create)
	group.Get("/", c.FindAll)
	group.Get("/:id", c.FindById)
	group.Put("/:id", c.Update)
	group.Delete("/:id", c.Delete)
}
