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
	if err := database.DB.Find(&posts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot fetch blog posts"})
	}
	return c.JSON(posts)
}

// Get a single blog post by ID
func GetBlogPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.BlogPost
	if err := database.DB.First(&post, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blog post not found"})
	}
	return c.JSON(post)
}

// Create a new blog post
func CreateBlogPost(c *fiber.Ctx) error {
	var post models.BlogPost
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	post.ID = uuid.New()
	if err := database.DB.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create blog post"})
	}
	return c.Status(fiber.StatusCreated).JSON(post)
}

// Update an existing blog post
func UpdateBlogPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.BlogPost
	if err := database.DB.First(&post, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blog post not found"})
	}

	var input models.BlogPost
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	post.Title = input.Title
	post.Content = input.Content
	post.Author = input.Author

	if err := database.DB.Save(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update blog post"})
	}

	return c.JSON(post)
}

// Delete a blog post
func DeleteBlogPost(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.BlogPost{}, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete blog post"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
