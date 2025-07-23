package main

import (
	"RAM/config"
	"RAM/middleware"
	"RAM/models"
	"RAM/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.LoadConfig()

	gormDB, err := models.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to init GORM DB: %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("Failed to get *sql.DB: %v", err)
	}
	defer sqlDB.Close()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Use(middleware.LoggingMiddleware)

	// Setup your routes here
	routes.SetupAuthRoutes(app, gormDB)
	routes.SetupDashboardRoutes(app, sqlDB)
	routes.SetupKeuanganRoutes(app, sqlDB)
	routes.SetupModalRoutes(app, sqlDB)
	routes.SetupKeuntunganRoutes(app, sqlDB)
	routes.SetupSusutRoutes(app, sqlDB)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
