package routes

import (
	"RAM/models"
	"RAM/payload"
	"RAM/middleware"
	"github.com/gofiber/fiber/v2"
	"database/sql"
)

func SetupKeuntunganRoutes(app *fiber.App, db *sql.DB) {
	app.Post("/keuntungan",middleware.JWTProtected(), func(c *fiber.Ctx) error {
		var req payload.EstimasiKeuntunganRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		e := models.EstimasiKeuntungan{
			SpCairPabrik:   req.SpCairPabrik,
			TonasiSpRam:    req.TonasiSpRam,
			HargaTbsBeliRam: req.HargaTbsBeliRam,
		}
		if err := models.CreateEstimasiKeuntungan(db, &e); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create profit estimate: " + err.Error()})
		}
		return c.JSON(fiber.Map{
			"message": "Profit estimate created and recorded as income",
			"data":    e,
		})
	})

	app.Get("/keuntungan", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		data, err := models.GetEstimasiKeuntungan(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		totalKeuntungan := 0
		for _, item := range data {
			totalKeuntungan += item.EstimasiKeuntungan
		}
		return c.JSON(fiber.Map{
			"data":            data,
			"total_keuntungan": totalKeuntungan,
		})
	})
}
