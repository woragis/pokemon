package controllers

import (
	"pokemon/database"
	"pokemon/models"
	"pokemon/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type AuthInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsernameInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var input AuthInput
	if err := c.BodyParser(&input); err != nil {
		log.Debug("BodyParser error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Debug("Password hash error: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	user := models.User{Username: input.Username, Email: input.Email, Password: hashed}
	if err := database.DB.Create(&user).Error; err != nil {
		log.Debug("DB insertion error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User exists or bad input"})
	}

	return c.JSON(fiber.Map{"message": "Registration successful"})
}

func UsernameLogin(c *fiber.Ctx) error {
	var input UsernameInput
	if err := c.BodyParser(&input); err != nil {
		log.Debug("BodyParser error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		log.Debug("DB query error:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		log.Debug("Password check error, invalid credentials")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		log.Debug("Password error: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": token})
}

func EmailLogin(c *fiber.Ctx) error {
	var input EmailInput
	if err := c.BodyParser(&input); err != nil {
		log.Debug("BodyParser error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		log.Debug("DB query error:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		log.Debug("Password check error, invalid credentials")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		log.Debug("Password error: ", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": token})
}

func Profile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)

	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.JSON(fiber.Map{"user":user})
}
