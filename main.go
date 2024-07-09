package main

import (
	"openmusic-api/exception"
	"openmusic-api/helper"
	"openmusic-api/model/web"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ResponseWithData{
			BaseResponse: web.BaseResponse{
				Code:   fiber.StatusInternalServerError,
				Status: "dwd",
			},
		})
	})

	err := app.Listen(":3000")
	helper.PanicIfError(err)
}
