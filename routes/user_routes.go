package routes

import (
	"go-fiber-api/handlers"
	"go-fiber-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/users", middlewares.JWTProtected())

	user.Get("/", handlers.GetUsers)
	user.Get("/:id", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
	user.Put("/:id", handlers.UpdateUser)
	user.Delete("/:id", handlers.DeleteUser)
}
