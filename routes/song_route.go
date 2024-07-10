package routes

import (
	"openmusic-api/controller"

	"github.com/gofiber/fiber/v2"
)

func SongRoutes(app *fiber.App, c controller.SongController) {
	group := app.Group("/songs")

	group.Post("/", c.Create)
	group.Get("/", c.FindAll)
	group.Get("/:id", c.FindById)
	group.Put("/:id", c.Update)
	group.Delete("/:id", c.Delete)
}
