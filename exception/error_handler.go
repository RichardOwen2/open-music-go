package exception

import (
	"openmusic-api/model/web"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).JSON(web.NewResponseWithMessage(
			e.Code,
			"error",
			e.Message,
		))
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.NewResponseWithMessage(
			fiber.StatusInternalServerError,
			"error",
			err.Error(),
		))
	}

	return nil
}
