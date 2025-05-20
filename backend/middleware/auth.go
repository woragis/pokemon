package middleware

import (
	"pokemon/database"
	"pokemon/models"
	"pokemon/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

func RequireRole(roles ...string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID := c.Locals("user_id").(uuid.UUID)
        var user models.User
        if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
            return c.SendStatus(fiber.StatusUnauthorized)
        }

        for _, r := range roles {
            if user.Role == r {
                return c.Next()
            }
        }

        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": "Insufficient permissions",
        })
    }
}
