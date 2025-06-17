package snapdex

import (
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type snapCommentHandler struct {
	service snapCommentService
}

func newSnapCommentHandler(service snapCommentService) *snapCommentHandler {
	return &snapCommentHandler{service: service}
}

func (h *snapCommentHandler) create(c *fiber.Ctx) error {
	var comment SnapComment
	if err := c.BodyParser(&comment); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}
	comment.UserID = userID

	if err := comment.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.service.create(&comment); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(comment)
}

func (h *snapCommentHandler) listByUser(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	limit, offset := utils.ParsePagination(c)
	comments, err := h.service.listByUser(userID, limit, offset)
	if err != nil {
		return err
	}

	return c.JSON(comments)
}

func (h *snapCommentHandler) countByUser(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	count, err := h.service.countByUser(userID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"count": count})
}

func (h *snapCommentHandler) updateStatus(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	status := c.Params("status")
	if !isValidSnapStatus(status) {
		return fiber.NewError(fiber.StatusBadRequest, "invalid status")
	}

	if err := h.service.updateStatus(id, status); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *snapCommentHandler) delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.delete(id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *snapCommentHandler) exists(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	exists, err := h.service.exists(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"exists": exists})
}
