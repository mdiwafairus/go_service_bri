package handlers

import (
	"go-fiber-api/constants"
	"go-fiber-api/helpers"
	"go-fiber-api/services"

	"github.com/gofiber/fiber/v2"
)

type TransactionRequest struct {
	Mid              string `json:"mid"`
	Nik              string `json:"nik"`
	NamaPupuk        string `json:"nama_pupuk"`
	NamaKomoditas    string `json:"nama_komoditas"`
	KgBeli           int    `json:"kg_beli"`
	TotalRupiah      int    `json:"total_rupiah"`
	RefNum           int    `json:"ref_num"`
	TanggalTransaksi int    `json:"tanggal_transaksi"`
}

type ReversalTransactionRequest struct {
	Nik    string `json:"nik"`
	Mid    string `json:"mid"`
	RefNum int    `json:"refnum"`
}

func TransactionHandler(service *services.TransactionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req TransactionRequest

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

		if req.RefNum <= 0 || req.TotalRupiah <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    constants.StatusRefNumRupiahTidakValid,
				"message": constants.MsgRefNumRupiahTidakValid,
			})
		}

		response, err := service.TransactionServiceResponse(req.Nik, req.Mid, req.NamaPupuk, req.NamaKomoditas, req.KgBeli, req.TotalRupiah, req.RefNum, req.TanggalTransaksi)
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
			case *services.DuplicateTransactionError:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"code":    constants.StatusTransaksiDuplikat,
					"message": constants.MsgTransaksiDuplikat,
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

func ReversalTransactionHandler(service *services.TransactionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req ReversalTransactionRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    "400",
				"message": "Invalid request body",
			})
		}

		if err := helpers.ValidateNIK(req.Nik); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    constants.StatusNikTidakValid,
				"message": constants.MsgNikTidakValid,
			})
		}
		if req.RefNum <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    constants.StatusRefNumRupiahTidakValid,
				"message": constants.MsgRefNumRupiahTidakValid,
			})
		}
		response, err := service.ReversalTransactionService(req.Nik, req.Mid, req.RefNum)
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
			case *services.DuplicateTransactionError:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"code":    constants.StatusTransaksiDuplikat,
					"message": constants.MsgTransaksiDuplikat,
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
