package main

import (
	"context"
	"go-fiber-api/config"
	"go-fiber-api/handlers"
	"go-fiber-api/middlewares"
	"go-fiber-api/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
	})

	config.InitLogger()

	log.Info().Msg("Server started on port 3000")
	log.Warn().Msg("This is a warning")

	config.ConnectDB()

	// middleware untuk cek database up/down
	app.Use(middlewares.DBCheckMiddleware)

	// global rate limiter
	app.Use(middlewares.RateLimiter())

	// check connection database
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 1*time.Second)
		defer cancel()

		if err := config.CheckDBHealth(ctx); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"status": "unhealthy",
				"error":  err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"status": "healthy",
		})
	})

	app.Post("/login", handlers.Login)

	routes.UserRoutes(app)
	routes.AllocationRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
