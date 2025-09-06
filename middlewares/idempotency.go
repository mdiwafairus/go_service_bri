package middlewares

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func IdempotencyMiddleware(rdb *redis.Client, ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Get("Idempotency-Key")
		if key == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Idempotency-Key required")
		}

		val, err := rdb.Get(ctx, key).Result()
		if err == nil && val != "" {
			c.Set("X-Idempotency", "HIT")
			return c.Status(fiber.StatusOK).SendString(val)
		}

		if err := c.Next(); err != nil {
			return err
		}

		body := string(c.Response().Body())
		if body != "" {
			_ = rdb.Set(ctx, key, body, 2*time.Minute).Err()
		}
		c.Set("X-Idempotency", "MISS")

		return nil
	}
}
