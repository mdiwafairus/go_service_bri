package routes

import (
	"go-fiber-api/config"
	"go-fiber-api/handlers"
	"go-fiber-api/middlewares"
	"go-fiber-api/repositories"
	"go-fiber-api/services"

	"github.com/gofiber/fiber/v2"
)

func AllocationRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	allocation := v1.Group("/allocation", middlewares.JWTProtected())

	allocationRepo := repositories.NewAllocationRepository(config.DB)
	allocationService := services.NewAllocationService(allocationRepo)

	allocation.Get("/quota", handlers.QuotaHandler(allocationService))
	allocation.Get("/inquiry", handlers.InquiryHandler(allocationService))
}
