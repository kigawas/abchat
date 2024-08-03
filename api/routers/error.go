package routers

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func handleError(c fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
}
