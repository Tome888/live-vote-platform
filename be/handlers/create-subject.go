package handlers

import (
	"my-fiber-app/structs"

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
)

func CreateSubjectHandler(db *sqlx.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req structs.CreateSubjectRequest

		rawRoomID := c.Locals("admin_room_id")

		var roomID int64
		if f, ok := rawRoomID.(float64); ok {
			roomID = int64(f)
		} else if i, ok := rawRoomID.(int64); ok {
			roomID = i
		}

		if err := c.Bind().JSON(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "flag": "1"})
		}

		var exists bool
		err := db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM room WHERE id = ?)", roomID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "DB check failed", "flag": "4"})
		}

		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "room not found", "flag": "5"})
		}

		res, err := db.Exec("INSERT INTO subjects (name, room_id) VALUES (?, ?)", req.Name, roomID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "flag": "6"})
		}

		subjectId, _ := res.LastInsertId()

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":    "subject created",
			"room_id":    roomID,
			"subject_id": subjectId,
			"name":       req.Name,
		})
	}
}
