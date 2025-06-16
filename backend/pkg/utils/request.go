package utils

import "github.com/gofiber/fiber/v2"

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
