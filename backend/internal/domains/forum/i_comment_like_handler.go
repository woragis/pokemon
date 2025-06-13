package forum

import (
	"net/http"
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *handler) createLike(c *fiber.Ctx) error {
	var like TopicCommentLike
	if err := c.BodyParser(&like); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	like.ID = userID
	if err := like.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.commentLikeService.create(&like); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(like)
}

func (h *handler) updateLike(c *fiber.Ctx) error {
	var like TopicCommentLike
	if err := c.BodyParser(&like); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	like.ID = userID
	if err := like.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.commentLikeService.update(&like); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(like)
}

func (h *handler) getLikes(c *fiber.Ctx) error {
	commentID, err := uuid.Parse(c.Params("comment_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid comment ID"})
	}
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}
	like, err := h.commentLikeService.get(commentID, userID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "like not found"})
	}
	return c.JSON(like)
}

func (h *handler) countLikes(c *fiber.Ctx) error {
	commentID, err := uuid.Parse(c.Params("comment_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid comment ID",
		})
	}

	likes, dislikes, err := h.commentLikeService.count(commentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"likes":    likes,
		"dislikes": dislikes,
	})
}

func (h *handler) deleteLike(c *fiber.Ctx) error {
	commentID, err := uuid.Parse(c.Params("comment_id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid comment ID"})
	}
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.commentLikeService.delete(commentID, userID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(http.StatusNoContent)
}
