package middleware

import (
	"pokemon/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if !strings.HasPrefix(authHeader, "Bearer ") {
            return c.SendStatus(fiber.StatusUnauthorized)
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := utils.ParseJWT(tokenStr)
        if err != nil || !token.Valid {
            return c.SendStatus(fiber.StatusUnauthorized)
        }

        claims := token.Claims.(jwt.MapClaims)
        c.Locals("user_id", uint(claims["user_id"].(float64)))
        return c.Next()
    }
}
