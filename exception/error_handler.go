package exception

import (
	"openmusic-api/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).JSON(web.NewResponseWithMessage(
			e.Code,
			"fail",
			e.Message,
		))
	}

	if e, ok := err.(validator.ValidationErrors); ok {
		return c.Status(fiber.StatusBadRequest).JSON(web.NewResponseWithMessage(
			fiber.StatusBadRequest,
			"fail",
			e.Error(),
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
