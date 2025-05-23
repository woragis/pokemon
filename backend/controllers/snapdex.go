package controllers

import (
	"pokemon/database"
	"pokemon/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GET /snaps
func GetSnaps(c *fiber.Ctx) error {
	var snaps []models.Snap
	if err :=database.DB.Find(&snaps).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch snaps"})
	}
	return c.JSON(snaps)
}

// GET /snaps/:id
func GetSnap(c *fiber.Ctx) error {
	id := c.Params("id")
	var snap models.Snap
	if err :=database.DB.First(&snap, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Snap not found"})
	}
	return c.JSON(snap)
}

// GET /snaps/me
func GetOwnSnaps(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	var snaps []models.Snap
	err := database.DB.Where("user_id = ?", userID).Find(&snaps).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}
	return c.JSON(snaps)
}

// POST /snaps
func CreateSnap(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var snap models.Snap
	if err := c.BodyParser(&snap); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	snap.ID = uuid.New()
	snap.UserID = userID

	if err := database.DB.Create(&snap).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create snap"})
	}

	return c.JSON(snap)
}

// DELETE /snaps/:id
func DeleteSnap(c *fiber.Ctx) error {
	id := c.Params("id")
	snapID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid snap ID"})
	}

	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Optional: only allow deleting own snap
	var snap models.Snap
	if err := database.DB.First(&snap, "id = ?", snapID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Snap not found"})
	}

	if snap.UserID != userID {
		return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
	}

	if err := database.DB.Delete(&snap).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete snap"})
	}

	return c.SendStatus(204)
}
