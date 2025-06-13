package middleware

import (
	"pokemon/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthRequired() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Authorization header required")
        }
        
        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid token")
        }
        
        c.Locals("userID", claims.UserID)
        c.Locals("email", claims.Email)
        
        return c.Next()
    }
}

func AuthOptional() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Next()
        }
        
        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            return c.Next()
        }
        
        c.Locals("userID", claims.UserID)
        c.Locals("email", claims.Email)
        
        return c.Next()
    }
}
