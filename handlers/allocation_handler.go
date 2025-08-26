package handlers

import (
	"go-fiber-api/constants"
	"go-fiber-api/services"

	"github.com/gofiber/fiber/v2"
)

type NikMidRequest struct {
	Mid string `json:"mid"`
	Nik string `json:"nik"`
}

func GetNikExistsHandler(service *services.AllocationService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var req NikMidRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    "400",
				"message": "Invalid request body",
			})
		}

		if req.Nik == "" || req.Mid == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    "400",
				"message": "nik dan mid wajib diisi",
			})
		}

		response, err := service.GetNikExistsResponse(req.Nik, req.Mid)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    constants.StatusNikNotFound,
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"responseCode":    constants.StatusSuccess,
			"responseMessage": constants.MsgSuccess,
			"data":            response,
		})
	}
}
