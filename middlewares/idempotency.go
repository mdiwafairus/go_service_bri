package middlewares

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func makeRequestHash(c *fiber.Ctx) string {
	raw := fmt.Sprintf("%s|%s|%s|%s",
		c.Method(),
		c.OriginalURL(),
		string(c.Body()),
		c.Get("Authorization"),
	)
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}

func IdempotencyMiddleware(rdb *redis.Client, ctx context.Context, shortTTL, longTTL time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqHash := makeRequestHash(c)

		shortKey := "idem:short:" + reqHash
		longKey := "idem:long:" + reqHash

		// short TTL check double click
		shortExists, _ := rdb.Exists(ctx, shortKey).Result()
		if shortExists > 0 {
			log.Warn().
				Str("path", c.OriginalURL()).
				Str("method", c.Method()).
				Str("reqHash", reqHash).
				Msg("Duplicate request blocked (double click)")
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Duplicate request detected (double click)",
			})
		}

		// long TTL (replay)
		longExists, _ := rdb.Exists(ctx, longKey).Result()
		if longExists > 0 {
			log.Warn().
				Str("path", c.OriginalURL()).
				Str("method", c.Method()).
				Str("reqHash", reqHash).
				Msg("Duplicate request blocked (replay)")
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Duplicate request detected (replay)",
			})
		}

		if err := rdb.Set(ctx, shortKey, 1, shortTTL).Err(); err != nil {
			log.Error().
				Err(err).
				Str("reqHash", reqHash).
				Msg("Failed to set short idempotency key")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to set short idempotency key",
			})
		}
		if err := rdb.Set(ctx, longKey, 1, longTTL).Err(); err != nil {
			log.Error().
				Err(err).
				Str("reqHash", reqHash).
				Msg("Failed to set long idempotency key")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to set long idempotency key",
			})
		}

		log.Info().
			Str("path", c.OriginalURL()).
			Str("method", c.Method()).
			Str("reqHash", reqHash).
			Dur("shortTTL", shortTTL).
			Dur("longTTL", longTTL).
			Msg("First request stored in Redis (idempotency keys set)")

		return c.Next()
	}
}
