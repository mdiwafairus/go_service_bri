package handlers

import (
	"go-fiber-api/constants"
	"go-fiber-api/dto"
	"go-fiber-api/helpers"
	"go-fiber-api/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

type QuotaRequest struct {
	Mid string `json:"mid"`
	Nik string `json:"nik"`
}

type InquiryRequest struct {
	Mid           string `json:"mid"`
	Nik           string `json:"nik"`
	NamaPupuk     string `json:"nama_pupuk"`
	NamaKomoditas string `json:"nama_komoditas"`
	KgBeli        int    `json:"kg_beli"`
}

func QuotaHandler(service *services.AllocationService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var req QuotaRequest

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

		response, err := service.QuotaServiceResponse(req.Nik, req.Mid)
		if err != nil {
			switch e := err.(type) {
			case *services.NikNotFoundError:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"code":    constants.StatusNikNotFound,
					"message": constants.MsgNikNotFound,
				})
			case *services.KiosNotMatchError:
				return c.JSON(dto.KiosTidakSesuaiResponse{
					Code:    constants.StatusStandUnsuitable,
					Message: constants.MsgStandUnsuitable,
					Suggest: e.Suggest,
				})
			default:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":    "500",
					"message": err.Error(),
				})
			}
		}

		return c.JSON(fiber.Map{
			"responseCode":    constants.StatusSuccess,
			"responseMessage": constants.MsgSuccess,
			"data":            response,
		})
	}
}

func InquiryHandler(service *services.AllocationService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		time.Sleep(2 * time.Second)

		var req InquiryRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    "400",
				"message": "Invalid request body",
			})
		}

		if !helpers.IsValidKomoditas(req.NamaKomoditas) || !helpers.IsValidPupuk(req.NamaPupuk) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    constants.StatusPupukKomoditasTidakValid,
				"message": constants.MsgPupukKomoditasTidakValid,
			})
		}

		if req.KgBeli <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    constants.StatusTidakMemilikiKuota,
				"message": constants.MsgTidakMemilikiKuota,
			})
		}

		response, err := service.InquiryServiceResponse(req.Nik, req.NamaKomoditas, req.Mid, req.NamaPupuk, req.KgBeli)
		if err != nil {
			switch err.(type) {
			case *services.NikNotFoundError:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"code":    constants.StatusNikNotFound,
					"message": constants.MsgNikNotFound,
				})
			case *services.KiosNotMatchError:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"code":    constants.StatusStandUnsuitable,
					"message": constants.MsgStandUnsuitable,
				})
			case *services.AllocationNotFound:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"code":    constants.StatusAlokasiNotFound,
					"message": constants.MsgAlokasiNotFound,
				})
			case *services.TidakMemilikiKuota:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"code":    constants.StatusTidakMemilikiKuota,
					"message": constants.MsgTidakMemilikiKuota,
				})
			case *services.PupukTidakValid:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"code":    constants.StatusPupukTidakValid,
					"message": constants.MsgPupukTidakValid,
				})
			default:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":    "500",
					"message": err.Error(),
				})
			}
		}

		return c.JSON(fiber.Map{
			"responseCode":    constants.StatusSuccess,
			"responseMessage": constants.MsgSuccess,
			"data":            response,
		})
	}
}
