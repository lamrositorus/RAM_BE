package routes

import (
	"RAM/models"
	"github.com/gofiber/fiber/v2"
	"database/sql"
)

func SetupDashboardRoutes(app *fiber.App, db *sql.DB) {
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		// Get data keuangan
		keuangan, err := models.GetKeuangan(db)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal ambil data keuangan"})
		}

		totalPemasukan, totalPengeluaran := 0, 0
		for _, k := range keuangan {
			if k.Tipe == "pemasukan" {
				totalPemasukan += k.Nominal
			} else {
				totalPengeluaran += k.Nominal
			}
		}

		// Get data modal
		modal, err := models.GetModal(db)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal ambil data modal"})
		}

		totalModal := 0
		for _, m := range modal {
			totalModal += m.TotalModal
		}

		// Get data keuntungan
		keuntungan, err := models.GetEstimasiKeuntungan(db)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal ambil data keuntungan"})
		}

		totalUntung := 0
		for _, k := range keuntungan {
			totalUntung += k.EstimasiKeuntungan
		}

		

		return c.JSON(fiber.Map{
			"total_pemasukan":   totalPemasukan,
			"total_pengeluaran": totalPengeluaran,
			"saldo_akhir":       totalPemasukan - totalPengeluaran,
			"total_modal":       totalModal,
			"total_untung":      totalUntung,			
		})
	})
}
