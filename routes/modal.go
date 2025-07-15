package routes

import (
	"RAM/models"
	"RAM/payload"
	"github.com/gofiber/fiber/v2"
	"database/sql"
	"RAM/middleware"
)

func SetupModalRoutes(app *fiber.App, db *sql.DB) {
app.Post("/modal",middleware.JWTProtected(), func(c *fiber.Ctx) error {
    var req payload.EstimasiModalRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
    }

    // Calculate total modal
    totalModal := req.SisaModal + (req.TonasiGantung * req.HargaTbs)

    modal := models.EstimasiModalRequest{
        SisaModal:     req.SisaModal,
        TonasiGantung: req.TonasiGantung,
        HargaTbs:      req.HargaTbs,
        TotalModal:    totalModal, // Set the calculated total modal here
    }

    if err := models.CreateModal(db, &modal); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(modal) // This will now include the correct TotalModal value
})


	app.Get("/modal",middleware.JWTProtected(), func(c *fiber.Ctx) error {
		modal, err := models.GetModal(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		totalModal := 0
		for _, m := range modal {
			totalModal += m.TotalModal
		}
		return c.JSON(fiber.Map{
			"data":        modal,
			"total_modal": totalModal,
		})
	})
}
