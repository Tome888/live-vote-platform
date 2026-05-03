package handlers

import (
	"fmt"
	"my-fiber-app/secret"
	"my-fiber-app/structs"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

func CreateSubToken(db *sqlx.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req structs.CreateSubTokenRequest
		secretKey := secret.GetSecret()
		if err := c.Bind().JSON(&req); err != nil {
			return err
		}

		parsedExpiry, err := time.Parse(time.RFC3339, req.ExpiresAt)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid expiration date format")
		}

		jwtContent := jwt.MapClaims{
			"room_id": req.RoomID,
			"role":    req.Role,
			"name":    req.Name,
			"iat":     time.Now().Unix(),
			"exp":     parsedExpiry.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtContent)
		finalToken, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		fmt.Println(req, "expires_at:", parsedExpiry)
		_, err = db.Exec("INSERT INTO tokens (room_id, voter_token, expires_at, name) VALUES (?, ?, ?, ?)",
			req.RoomID, finalToken, parsedExpiry, req.Name)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":    "Sub token created successfully",
			"token":      finalToken,
			"room_id":    req.RoomID,
			"role":       req.Role,
			"name":       req.Name,
			"expires_at": req.ExpiresAt,
		})
	}
}
