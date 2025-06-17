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

type snapLikeHandler struct {
	service snapLikeService
}

func newSnapLikeHandler(service snapLikeService) *snapLikeHandler {
	return &snapLikeHandler{service}
}

func (h *snapLikeHandler) like(c *fiber.Ctx) error {
	snapID, err := uuid.Parse(c.Params("snap_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid snap_id")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	like := &SnapLike{
		SnapID: snapID,
		UserID: userID,
	}

	if err := like.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.service.like(like); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *snapLikeHandler) unlike(c *fiber.Ctx) error {
	snapID, err := uuid.Parse(c.Params("snap_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid snap_id")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	if err := h.service.unlike(snapID, userID); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *snapLikeHandler) deleteAllBySnap(c *fiber.Ctx) error {
	snapID, err := uuid.Parse(c.Params("snap_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid snap_id")
	}

	if err := h.service.deleteAllBySnap(snapID); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *snapLikeHandler) listUserLikes(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	likes, err := h.service.listUserLikes(userID)
	if err != nil {
		return err
	}

	return c.JSON(likes)
}

func (h *snapLikeHandler) isLikedByUser(c *fiber.Ctx) error {
	snapID, err := uuid.Parse(c.Params("snap_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid snap_id")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	isLiked, err := h.service.isLikedByUser(snapID, userID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"liked": isLiked})
}
