package controllers

import (
	"pokemon/database"
	"pokemon/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllUsers(c *fiber.Ctx) error {
    var users []models.User
    if err := database.DB.Find(&users).Error; err != nil {
        return c.SendStatus(fiber.StatusInternalServerError)
    }
    return c.JSON(users)
}

func BanUser(c *fiber.Ctx) error {
    id := c.Params("id")
    return database.DB.Model(&models.User{}).
        Where("id = ?", id).
        Update("role", "banned").
        Error
}

func ModerateShouts(c *fiber.Ctx) error {
	var shouts []models.Shout
	if err := database.DB.Preload("User").Order("created_at DESC").Find(&shouts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch shouts",
		})
	}
	return c.JSON(shouts)
}

func DeleteShout(c *fiber.Ctx) error {
	idStr := c.Params("id")
	shoutID, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid shout ID"})
	}

	// Delete shout (and cascade if necessary)
	if err := database.DB.Delete(&models.Shout{}, "id = ?", shoutID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete shout",
		})
	}

	return c.JSON(fiber.Map{"message": "Shout deleted successfully"})
}
