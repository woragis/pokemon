package controllers

import (
	"time"

	"pokemon/database"
	"pokemon/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateForumTopic(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	var input struct {
		Title      string    `json:"title"`
		CategoryID uuid.UUID `json:"categoryId"`
		Pinned     bool      `json:"pinned"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	topic := models.ForumTopic{
		Title:      input.Title,
		AuthorID:   userID,
		CategoryID: input.CategoryID,
		Pinned:     input.Pinned,
		CreatedAt:  time.Now(),
	}
	if err := database.DB.Create(&topic).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(topic)
}

func GetForumTopics(c *fiber.Ctx) error {
	var topics []models.ForumTopic
	if err := database.DB.Preload("Author").Preload("Category").Find(&topics).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(topics)
}

func GetForumTopicById(c *fiber.Ctx) error {
	id := c.Params("id")
	topicID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	var topic models.ForumTopic
	if err := database.DB.Preload("Author").Preload("Category").First(&topic, "id = ?", topicID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(topic)
}

func EditForumTopicById(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	topicID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	var topic models.ForumTopic
	if err := database.DB.First(&topic, "id = ?", topicID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	if topic.AuthorID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
	}
	var input struct {
		Title      *string    `json:"title"`
		CategoryID *uuid.UUID `json:"categoryId"`
		Pinned     *bool      `json:"pinned"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if input.Title != nil {
		topic.Title = *input.Title
	}
	if input.CategoryID != nil {
		topic.CategoryID = *input.CategoryID
	}
	if input.Pinned != nil {
		topic.Pinned = *input.Pinned
	}
	topic.UpdatedAt = time.Now()
	if err := database.DB.Save(&topic).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(topic)
}

func DeleteForumTopicById(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	topicID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	var topic models.ForumTopic
	if err := database.DB.First(&topic, "id = ?", topicID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	if topic.AuthorID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
	}
	if err := database.DB.Delete(&topic).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func CreateForumCategory(c *fiber.Ctx) error {
	var input models.ForumCategory
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if err := database.DB.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(input)
}

func GetForumCategories(c *fiber.Ctx) error {
	var categories []models.ForumCategory
	if err := database.DB.Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(categories)
}

func GetForumCategoryById(c *fiber.Ctx) error {
	uid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	var category models.ForumCategory
	if err := database.DB.First(&category, "id = ?", uid).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(category)
}

func EditForumCategoryById(c *fiber.Ctx) error {
	uid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	var input models.ForumCategory
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	var category models.ForumCategory
	if err := database.DB.First(&category, "id = ?", uid).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	category.Name = input.Name
	category.Color = input.Color
	category.Description = input.Description
	if err := database.DB.Save(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(category)
}

func DeleteForumCategoryById(c *fiber.Ctx) error {
	uid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}
	if err := database.DB.Delete(&models.ForumCategory{}, "id = ?", uid).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
