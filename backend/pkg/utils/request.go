package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ParsePagination(c *fiber.Ctx) (limit, offset int) {
	limit = 20
	offset = 0

	if l := c.QueryInt("limit"); l > 0 {
		limit = l
	}
	if o := c.QueryInt("offset"); o >= 0 {
		offset = o
	}
	return
}

// ParseUUID parses a string into a UUID. If invalid, returns uuid.Nil (000...000).
func ParseUUID(id string) uuid.UUID {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil
	}
	return parsed
}
