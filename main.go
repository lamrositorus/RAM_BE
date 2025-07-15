package main

import (
	"RAM/config"
	"RAM/middleware"
	"RAM/routes"
	"RAM/utils"
	"log"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	// Init GORM DB
	gormDB, err := utils.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to init GORM DB: %v", err)
	}

	// Get sql.DB for raw SQL
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("Failed to get *sql.DB: %v", err)
	}

	app := fiber.New()
app.Use(cors.New(cors.Config{
    AllowOrigins: "*", // atau "http://localhost:5173" untuk lebih aman
    AllowHeaders: "Origin, Content-Type, Accept, Authorization",
}))

	app.Use(middleware.LoggingMiddleware)

	// Routes
	routes.SetupAuthRoutes(app, gormDB)       // pakai GORM
	routes.SetupDashboardRoutes(app, sqlDB)
	routes.SetupKeuanganRoutes(app, sqlDB)    // pakai sql.DB
	routes.SetupModalRoutes(app, sqlDB)     // bisa sama juga
	routes.SetupKeuntunganRoutes(app, sqlDB)
	routes.SetupSusutRoutes(app, sqlDB)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
