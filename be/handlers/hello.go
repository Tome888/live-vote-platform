package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func HelloHandler(c fiber.Ctx) error {
	fmt.Println("Voting Server is up and runing")
	return c.SendString("Voting Server is up and runing")
}
