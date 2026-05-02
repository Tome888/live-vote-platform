package handlers

import (
	"fmt"
	"my-fiber-app/secret"
	"my-fiber-app/structs"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

func CreateSubjectHandler(db *sqlx.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req structs.CreateSubjectRequest
		secretStr := secret.GetSecret()
		claims := &structs.RoomToken{}
		var exists bool

		if err := c.Bind().JSON(&req); err != nil {
			fmt.Println("1")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "flag": "1"})
		}

		if req.Token == "" {
			fmt.Println("2")

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "token is required", "flag": "2"})
		}

		cleanToken := strings.Trim(req.Token, " \n\r\t\"")

		_, err := jwt.ParseWithClaims(cleanToken, claims, func(token *jwt.Token) (any, error) {
			return []byte(secretStr), nil
		})

		if err != nil {
			fmt.Println("3", req.Token, claims)

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "invalid token",
				"details": err.Error(),
			})
		}

		err = db.QueryRow(`SELECT EXISTS(SELECT 1 FROM room WHERE id = ? AND name = ?)`, claims.RoomId, claims.Name).Scan(&exists)

		if err != nil {
			fmt.Println("4")

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "flag": "4"})
		}

		if !exists {
			fmt.Println("5")

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "room not found", "flag": "5"})
		}

		res, err := db.Exec("INSERT INTO subjects (name, room_id) VALUES (?, ?)", req.Name, claims.RoomId)

		if err != nil {
			fmt.Println("6")

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "flag": "6"})
		}

		subjectId, err := res.LastInsertId()

		if err != nil {
			fmt.Println("7")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "flag": "7"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "subject created", "room_id": claims.RoomId, "subject_id": subjectId, "name": req.Name})
	}
}
