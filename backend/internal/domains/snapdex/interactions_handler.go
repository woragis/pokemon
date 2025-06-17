package snapdex

import (
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *handler) createComment(c *fiber.Ctx) error {
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

	if err := h.iCommentS.create(&comment); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(comment)
}

func (h *handler) listCommentsByUser(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	limit, offset := utils.ParsePagination(c)
	comments, err := h.iCommentS.listByUser(userID, limit, offset)
	if err != nil {
		return err
	}

	return c.JSON(comments)
}

func (h *handler) countCommentsByUser(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	count, err := h.iCommentS.countByUser(userID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"count": count})
}

func (h *handler) updateCommentStatus(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	status := c.Params("status")
	if !isValidSnapStatus(status) {
		return fiber.NewError(fiber.StatusBadRequest, "invalid status")
	}

	if err := h.iCommentS.updateStatus(id, status); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *handler) deleteComment(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	if err := h.iCommentS.delete(id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *handler) commentExists(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	exists, err := h.iCommentS.exists(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"exists": exists})
}

func (h *handler) likeSnap(c *fiber.Ctx) error {
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

	if err := h.iLikeS.like(like); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) unlikeSnap(c *fiber.Ctx) error {
	snapID, err := uuid.Parse(c.Params("snap_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid snap_id")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	if err := h.iLikeS.unlike(snapID, userID); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *handler) deleteAllLikesBySnap(c *fiber.Ctx) error {
	snapID, err := uuid.Parse(c.Params("snap_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid snap_id")
	}

	if err := h.iLikeS.deleteAllBySnap(snapID); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *handler) listUserLikes(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	likes, err := h.iLikeS.listUserLikes(userID)
	if err != nil {
		return err
	}

	return c.JSON(likes)
}

func (h *handler) isLikedByUser(c *fiber.Ctx) error {
	snapID, err := uuid.Parse(c.Params("snap_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid snap_id")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	isLiked, err := h.iLikeS.isLikedByUser(snapID, userID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"liked": isLiked})
}

func (h *handler) likeComment(c *fiber.Ctx) error {
	commentID, err := uuid.Parse(c.Params("comment_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment_id")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	like := &SnapCommentLike{
		CommentID: commentID,
		UserID:    userID,
	}

	if err := like.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.iCommentLikeS.like(like); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) unlikeComment(c *fiber.Ctx) error {
	commentID, err := uuid.Parse(c.Params("comment_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment_id")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	if err := h.iCommentLikeS.unlike(commentID, userID); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *handler) listLikesByComment(c *fiber.Ctx) error {
	commentID, err := uuid.Parse(c.Params("comment_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment_id")
	}

	likes, err := h.iCommentLikeS.listByComment(commentID)
	if err != nil {
		return err
	}

	return c.JSON(likes)
}

func (h *handler) listCommentsUserLikes(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	likes, err := h.iCommentLikeS.listUserLikes(userID)
	if err != nil {
		return err
	}

	return c.JSON(likes)
}

func (h *handler) isCommentLikedByUser(c *fiber.Ctx) error {
	commentID, err := uuid.Parse(c.Params("comment_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment_id")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	liked, err := h.iCommentLikeS.isLikedByUser(commentID, userID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"liked": liked})
}
