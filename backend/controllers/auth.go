package controllers

import (
	"pokemon/database"
	"pokemon/models"
	"pokemon/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var input AuthInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	user := models.User{Username: input.Username, Email: input.Email, Password: hashed}
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User exists or bad input"})
	}

	return c.JSON(fiber.Map{"message": "Registration successful"})
}

func Login(c *fiber.Ctx) error {
	var input AuthInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": token})
}
