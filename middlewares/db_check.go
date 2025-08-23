package middlewares

import (
	"context"
	"go-fiber-api/config"
	"time"

	"github.com/gofiber/fiber/v2"
)

func DBCheckMiddleware(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 1*time.Second)
	defer cancel()

	if err := config.CheckDBHealth(ctx); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "Database is down or unresponsive",
		})
	}

	return c.Next()
}
