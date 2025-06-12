package forum

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *handler) createCategory(c *fiber.Ctx) error {
	var category TopicCategory
	if err := c.BodyParser(&category); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	if err := h.i_category_s.create(&category); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(category)
}

func (h *handler) getCategoryByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	category, err := h.i_category_s.getByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(category)
}

func (h *handler) updateCategory(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	var category TopicCategory
	if err := c.BodyParser(&category); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	category.ID = id
	if err := h.i_category_s.update(&category); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(category)
}

func (h *handler) deleteCategory(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	if err := h.i_category_s.delete(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(http.StatusNoContent)
}

func (h *handler) listCategories(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)
	categories, err := h.i_category_s.list(limit, offset)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(categories)
}
