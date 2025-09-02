package helpers

import (
	"github.com/gofiber/fiber/v2"
)

// ResponseSuccess untuk response sukses
func ResponseSuccess(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "00",
		"message": "Success",
		"data":    data,
	})
}

// ResponseError untuk response error
func ResponseError(c *fiber.Ctx, code string, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"code":    code,
		"message": message,
	})
}
