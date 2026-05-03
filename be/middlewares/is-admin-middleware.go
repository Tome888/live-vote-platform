package middlewares

import (
	"my-fiber-app/secret"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func IsAdminMiddleware(c fiber.Ctx) error {
	jwtSecret := secret.GetSecret()
	authHeader := c.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing or invalid token format")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return fiber.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.ErrUnauthorized
	}

	if claims["role"] != "admin" {
		return fiber.NewError(fiber.StatusForbidden, "Admin privileges required")
	}

	c.Locals("admin_room_id", claims["room_id"])
	c.Locals("user_role", claims["role"])

	return c.Next()
}
