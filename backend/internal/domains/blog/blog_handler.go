package blog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type handler struct {
	s postService
}

func NewHandler(db *gorm.DB, redis *redis.Client) *handler {
	repo := newRepository(db)
	service := newService(repo, redis)

	return &handler{
		s: service,
	}
}

func (h *handler) createPost(c *fiber.Ctx) error {
	var post Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if err := h.s.createPost(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(post)
}

func (h *handler) getPost(c *fiber.Ctx) error {
	id, err := uuid.Parse("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid post ID"})
	}
	post, err := h.s.getPost(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "post not found"})
	}
	return c.JSON(post)
}

func (h *handler) listPostsByUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	limit := c.QueryInt("limit", 10)
	if limit < 1 {
		limit = 10
	}
	offset := c.QueryInt("offset", 0)
	if offset < 0 {
		offset = 0
	}

	posts, err := h.s.listPostsByUser(userID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(posts)
}

func (h *handler) listPosts(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	if limit < 1 {
		limit = 10
	}
	offset := c.QueryInt("offset", 0)
	if offset < 0 {
		offset = 0
	}

	posts, err := h.s.listPosts(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(posts)
}

func (h *handler) updatePost(c *fiber.Ctx) error {
	id, err := uuid.Parse("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid post ID"})
	}

	var post Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	post.ID = id

	if err := h.s.updatePost(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(post)
}

func (h *handler) deletePost(c *fiber.Ctx) error {
	id, err := uuid.Parse("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid post ID"})
	}
	if err := h.s.deletePost(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
