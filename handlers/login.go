package handlers

import (
	"go-fiber-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {

	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.Username == "test" && req.Password == "123456" {
		token, _ := middlewares.GenerateJWT(1)
		return c.JSON(fiber.Map{"token": token})
	}

	return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
}
