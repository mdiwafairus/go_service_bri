package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRoute(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber! 🚀")
	})

	request := httptest.NewRequest("GET", "/", nil)
	response, _ := app.Test(request)

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", response.StatusCode)
	}

	body := make([]byte, response.ContentLength)
	response.Body.Read(body)
	if string(body) != "Hello, Fiber! 🚀" {
		t.Errorf("Expected body 'Hello, Fiber! 🚀', got %q", string(body))
	}
}
