package forum

import (
	"errors"
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *handler) countTopicLikes(c *fiber.Ctx) error {
	topicID, err := uuid.Parse(c.Params("topic_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid topic ID"})
	}

	likes, err := h.likeService.count(topicID, true)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	dislikes, err := h.likeService.count(topicID, false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"likes":    likes,
		"dislikes": dislikes,
	})
}

func (h *handler) createTopicLike(c *fiber.Ctx) error {
	topicID, err := uuid.Parse(c.Params("topic_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid topic ID"})
	}

	var input struct {
		Like bool `json:"like"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	like := &TopicLike{
		TopicID: topicID,
		UserID:  userID,
		Like:    input.Like,
	}

	if err := h.likeService.create(like); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) getTopicLike(c *fiber.Ctx) error {
	topicID, err := uuid.Parse(c.Params("topic_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid topic ID"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	like, err := h.likeService.get(topicID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(like)
}

func (h *handler) deleteTopicLike(c *fiber.Ctx) error {
	topicID, err := uuid.Parse(c.Params("topic_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid topic ID"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.likeService.delete(topicID, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
