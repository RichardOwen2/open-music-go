package helper

import "github.com/gofiber/fiber/v2"


func ErrorIfNotExist(message string, exist bool, err error) error {
	if err != nil {
		return err
	}

	if !exist {
		return fiber.NewError(fiber.StatusNotFound, message)
	}

	return nil
}
