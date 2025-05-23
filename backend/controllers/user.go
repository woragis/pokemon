package controllers

import (
	"pokemon/database"
	"pokemon/models"

	"github.com/gofiber/fiber/v2"
)

func GetUserRoles(c *fiber.Ctx) error {
	userID := c.Params("id")
	var user models.User

	if err := database.DB.Preload("Roles").First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user.Roles)
}

func AssignRoleToUser(c *fiber.Ctx) error {
	type Request struct {
		RoleName string `json:"role_name"`
	}
	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	userID := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var role models.Role
	if err := database.DB.Where("name = ?", body.RoleName).First(&role).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}

	if err := database.DB.Model(&user).Association("Roles").Append(&role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not assign role"})
	}

	return c.JSON(fiber.Map{"message": "Role assigned successfully"})
}