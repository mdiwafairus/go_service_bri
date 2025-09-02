package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        10,              // max 10 req
		Expiration: 1 * time.Minute, // reset setiap 1 menit

		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Too many requests, please try again later"})
		},
	})
}
