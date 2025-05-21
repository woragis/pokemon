package controllers

import (
	"pokemon/database"
	"pokemon/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Get all blog posts
func GetBlogPosts(c *fiber.Ctx) error {
	var posts []models.BlogPost
	if err := database.DB.Preload("Author").Find(&posts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot fetch blog posts"})
	}
	return c.JSON(fiber.Map{"posts": posts})
}

// Get a single blog post by ID
func GetBlogPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.BlogPost
	if err := database.DB.Preload("Author").First(&post, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blog post not found"})
	}
	return c.JSON(fiber.Map{"post": post})
}

// Create a new blog post
func CreateBlogPost(c *fiber.Ctx) error {
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	authorID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized or missing user ID"})
	}

	post := models.BlogPost{
		ID:       uuid.New(),
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: authorID,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create blog post"})
	}

	// Preload author before returning
	database.DB.Preload("Author").First(&post, "id = ?", post.ID)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"post": post})
}

// Update an existing blog post
func UpdateBlogPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.BlogPost
	if err := database.DB.First(&post, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blog post not found"})
	}

	// Optional: check that the requesting user is the author
	if userID, ok := c.Locals("user_id").(uuid.UUID); !ok || userID != post.AuthorID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not the author of this post"})
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	post.Title = input.Title
	post.Content = input.Content

	if err := database.DB.Save(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update blog post"})
	}

	database.DB.Preload("Author").First(&post, "id = ?", post.ID)

	return c.JSON(fiber.Map{"post": post})
}

// Delete a blog post
func DeleteBlogPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.BlogPost
	if err := database.DB.First(&post, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blog post not found"})
	}

	if userID, ok := c.Locals("user_id").(uuid.UUID); !ok || userID != post.AuthorID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not the author of this post"})
	}

	if err := database.DB.Delete(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete blog post"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
