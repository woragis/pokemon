package controllers

import (
	"pokemon/database"
	"pokemon/models"

	"github.com/gofiber/fiber/v2"
)

func PostShout(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var body struct {
		Content string `json:"content"`
	}
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	shout := models.Shout{
		UserID:  userID,
		Content: body.Content,
	}

	if err := database.DB.Create(&shout).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(shout)
}

func GetShoutFeed(c *fiber.Ctx) error {
	var shouts []models.Shout
	if err := database.DB.Preload("User").Order("created_at desc").Limit(50).Find(&shouts).Error; err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(shouts)
}

func LikeShout(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	shoutID := c.Params("id")

	like := models.ShoutLike{
		UserID:  userID,
		ShoutID: parseUint(shoutID),
	}

	if err := database.DB.Create(&like).Error; err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{"message": "liked"})
}
