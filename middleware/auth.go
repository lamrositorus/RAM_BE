package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	log.Printf("Request: %s %s took %v", c.Method(), c.Path(), time.Since(start))
	return err
}
