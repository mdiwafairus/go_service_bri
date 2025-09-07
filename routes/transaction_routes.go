package routes

import (
	"go-fiber-api/config"
	"go-fiber-api/handlers"
	"go-fiber-api/middlewares"
	"go-fiber-api/repositories"
	"go-fiber-api/services"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	transaction := v1.Group("/transaction", middlewares.JWTProtected(), middlewares.IdempotencyMiddleware(config.RedisCli, config.RedisCtx, config.ShortTTL, config.LongTTL))

	transactionRepo := repositories.NewTransactionRepository(config.DB)
	transactionService := services.NewTransactionService(transactionRepo)

	transaction.Get("/purchase", handlers.TransactionHandler(transactionService))
}
