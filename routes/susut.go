package routes

import (
	"RAM/models"
	"RAM/payload"
	"github.com/gofiber/fiber/v2"
	"database/sql"
	"RAM/middleware"
)

func SetupSusutRoutes(app *fiber.App, db *sql.DB) {
	app.Post("/susut",middleware.JWTProtected(), func(c *fiber.Ctx) error {
		var req payload.SusutTimbanganRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request: " + err.Error()})
		}
		susut := models.SusutTimbangan{
			NomorPolisi:    req.NomorPolisi,
			NamaSupir:      req.NamaSupir,
			JamMasukPabrik: req.JamMasukPabrik,
			SPPabrik:       req.SPPabrik,
			BuahPulangan:   req.BuahPulangan,
			JamKeluarRAM:   req.JamKeluarRAM,
			SPRAM:          req.SPRAM,
		}
		if err := models.CreateSusutTimbangan(db, &susut); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error: " + err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(susut)
	})

	app.Get("/susut",middleware.JWTProtected(), func(c *fiber.Ctx) error {
		data, err := models.GetSusutTimbangan(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		var totalKelebihan, totalSusut float64
		for _, item := range data {
			if item.Status == "kelebihan" {
				totalKelebihan += item.Selisih
			} else if item.Status == "susut" {
				totalSusut += -item.Selisih
			}
		}
		return c.JSON(fiber.Map{
			"data": data,
			"summary": fiber.Map{
				"total_kelebihan": totalKelebihan,
				"total_susut":     totalSusut,
				"net_balance":     totalKelebihan - totalSusut,
			},
		})
	})
}
