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

func RequireAuth() fiber.Handler {
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

		userIDStr := claims["user_id"].(string)
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Locals("user_id", userID)
		return c.Next()
	}
}

func RequireRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		var user models.User
		if err := database.DB.Preload("Roles").First(&user, "id = ?", userID).Error; err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Check if user has at least one of the required roles
		for _, userRole := range user.Roles {
			for _, requiredRole := range roles {
				if userRole.Name == requiredRole {
					return c.Next()
				}
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Insufficient permissions",
		})
	}
}

func RequireAllRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(uuid.UUID)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		var user models.User
		if err := database.DB.Preload("Roles").First(&user, "id = ?", userID).Error; err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		roleSet := make(map[string]bool)
		for _, r := range user.Roles {
			roleSet[r.Name] = true
		}

		for _, required := range roles {
			if !roleSet[required] {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "Missing required role: " + required,
				})
			}
		}

		return c.Next()
	}
}
