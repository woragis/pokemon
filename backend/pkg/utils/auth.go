package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// helper to get userID from locals
func GetUserIDFromLocals(c *fiber.Ctx) (uuid.UUID, error) {
	idVal := c.Locals("userID")
	if idVal == nil {
		return uuid.Nil, errors.New("user ID not found in context")
	}

	switch v := idVal.(type) {
	case uuid.UUID:
		return v, nil
	case string:
		return uuid.Parse(v)
	default:
		return uuid.Nil, errors.New("invalid user ID type in context")
	}
}
