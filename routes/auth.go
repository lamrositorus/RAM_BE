package routes

import (
	"RAM/middleware"
	"RAM/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SetupAuthRoutes(app *fiber.App, db *gorm.DB) {
	auth := app.Group("/auth")

	auth.Post("/signup", func(c *fiber.Ctx) error {
		var input models.User
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		input.Password = string(hashedPassword)

		if err := db.Create(&input).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "User creation failed"})
		}

		return c.JSON(fiber.Map{"message": "Signup successful"})
	})

	auth.Post("/signin", func(c *fiber.Ctx) error {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		var user models.User
		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Email not found"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Incorrect password"})
		}

		token, err := middleware.GenerateToken(user.ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Could not create token"})
		}

		return c.JSON(fiber.Map{"token": token})
	})
}