package handlers

import (
	"my-fiber-app/secret"
	"my-fiber-app/structs"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

func CreateRoomHandler(db *sqlx.DB) fiber.Handler {
	return func(c fiber.Ctx) error {

		var roomData structs.CreateRoomRequest
		secretStr := secret.GetSecret()

		err := c.Bind().Body(&roomData)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		res, err := db.Exec("INSERT INTO room (name) VALUES (?)", roomData.Name)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		roomID, err := res.LastInsertId()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Could not get room ID")
		}

		jwtContent := jwt.MapClaims{
			"room_id": roomID,
			"name":    roomData.Name,
			"iat":     time.Now().Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtContent)
		connectionKey, err := token.SignedString([]byte(secretStr))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Could not generate key")
		}

		_, err = db.Exec("UPDATE room SET connection_key = ? WHERE id = ?", connectionKey, roomID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(201).JSON(fiber.Map{
			"info":           jwtContent,
			"connection_key": connectionKey,
		})
	}
}
