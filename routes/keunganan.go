package routes

import (
	"RAM/middleware"
	"RAM/models"
	"RAM/payload"
	"github.com/gofiber/fiber/v2"
	"database/sql"
)

func SetupKeuanganRoutes(app *fiber.App, db *sql.DB) {
	app.Post("/keuangan", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		var req payload.KeuanganRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		keuangan := models.Keuangan{
			Deskripsi: req.Deskripsi,
			Nominal:   req.Nominal,
			Tipe:      req.Tipe,
		}

		if err := models.CreateKeuangan(db, &keuangan); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(keuangan)
	})

	app.Get("/keuangan", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		keuangan, err := models.GetKeuangan(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		totalPemasukan, totalPengeluaran := 0, 0
		for _, k := range keuangan {
			if k.Tipe == "pemasukan" {
				totalPemasukan += k.Nominal
			} else {
				totalPengeluaran += k.Nominal
			}
		}

		return c.JSON(fiber.Map{
			"data":             keuangan,
			"total_pemasukan":  totalPemasukan,
			"total_pengeluaran": totalPengeluaran,
			"saldo_akhir":      totalPemasukan - totalPengeluaran,
		})
	})
}