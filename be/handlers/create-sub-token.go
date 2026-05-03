package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

func CreateSubToken(db *sqlx.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Sub token created successfully",
		})
	}
}
