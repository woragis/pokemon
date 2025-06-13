package forum

import (
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *handler) createView(c *fiber.Ctx) error {
	topicID, err := uuid.Parse(c.Params("topic_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid topic ID"})
	}

	var view TopicView
	if err := c.BodyParser(&view); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	view.TopicID = topicID
	view.IPAddress = c.IP()

	// Check if user_id is present in locals (i.e., user is authenticated)
	if uid := c.Locals("user_id"); uid != nil {
		if userID, ok := uid.(uuid.UUID); ok {
			view.UserID = userID
		}
	}

	if err := h.viewService.create(&view); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) listViewsByUser(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	views, err := h.viewService.listByUser(userID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(views)
}
