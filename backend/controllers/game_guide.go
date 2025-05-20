package controllers

import (
	"pokemon/database"
	"pokemon/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateGameGuide(c *fiber.Ctx) error {
    userID := c.Locals("user_id").(uuid.UUID)

    var input struct {
        Title   string   `json:"title"`
        Slug    string   `json:"slug"`
        Summary string   `json:"summary"`
        Content string   `json:"content"`
        Tags    []string `json:"tags"`
    }

    if err := c.BodyParser(&input); err != nil || input.Title == "" || input.Slug == "" || input.Content == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing fields"})
    }

    guide := models.GameGuide{
        ID:        uuid.New(),
        Title:     input.Title,
        Slug:      input.Slug,
        Summary:   input.Summary,
        Content:   input.Content,
        AuthorID:  userID,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    // Attach tags
    for _, name := range input.Tags {
        tag := models.GameGuideTag{}
        database.DB.FirstOrCreate(&tag, models.GameGuideTag{Name: name})
        guide.Tags = append(guide.Tags, tag)
    }

    if err := database.DB.Create(&guide).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create guide"})
    }

    return c.JSON(guide)
}

func ListGameGuides(c *fiber.Ctx) error {
    var guides []models.GameGuide
    if err := database.DB.Preload("Tags").Order("created_at desc").Find(&guides).Error; err != nil {
        return c.SendStatus(fiber.StatusInternalServerError)
    }
    return c.JSON(guides)
}

func GetGameGuide(c *fiber.Ctx) error {
    slug := c.Params("slug")
    var guide models.GameGuide
    if err := database.DB.Preload("Tags").First(&guide, "slug = ?", slug).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Guide not found"})
    }
    return c.JSON(guide)
}

func UpdateGameGuide(c *fiber.Ctx) error {
    slug := c.Params("slug")
    var guide models.GameGuide
    if err := database.DB.First(&guide, "slug = ?", slug).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Guide not found"})
    }

    var input struct {
        Title   string   `json:"title"`
        Summary string   `json:"summary"`
        Content string   `json:"content"`
        Tags    []string `json:"tags"`
    }
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    guide.Title = input.Title
    guide.Summary = input.Summary
    guide.Content = input.Content
    guide.UpdatedAt = time.Now()

    if len(input.Tags) > 0 {
        guide.Tags = nil
        for _, name := range input.Tags {
            tag := models.GameGuideTag{}
            database.DB.FirstOrCreate(&tag, models.GameGuideTag{Name: name})
            guide.Tags = append(guide.Tags, tag)
        }
    }

    if err := database.DB.Save(&guide).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update guide"})
    }

    return c.JSON(guide)
}
