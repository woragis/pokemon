package shout

import (
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/********
 * LIKE *
 ********/

func (h *handler) createLike(c *fiber.Ctx) error {
	shoutIDParam := c.Params("shout_id")
	shoutID, err := uuid.Parse(shoutIDParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid shout ID")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	var like ShoutLike
	if err := c.BodyParser(&like); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	like.ShoutID = shoutID
	like.UserID = userID

	if err := like.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.is.createLike(&like); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) deleteLike(c *fiber.Ctx) error {
	shoutID, err := uuid.Parse(c.Params("shout_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid shout ID")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	if err := h.is.deleteLike(shoutID, userID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

/************
 * COMMENTS *
 ************/

func (h *handler) createComment(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	var comment ShoutComment
	if err := c.BodyParser(&comment); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	comment.UserID = userID

	if err := comment.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.is.createComment(&comment); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(comment)
}

func (h *handler) updateComment(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	var comment ShoutComment
	if err := c.BodyParser(&comment); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	comment.UserID = userID

	if err := comment.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.is.updateComment(&comment); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(comment)
}

func (h *handler) deleteComment(c *fiber.Ctx) error {
	commentID, err := uuid.Parse(c.Params("comment_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid comment ID")
	}

	if err := h.is.deleteComment(commentID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

/********
 * VIEW *
 ********/

func (h *handler) createView(c *fiber.Ctx) error {
	shoutID, err := uuid.Parse(c.Params("shout_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid shout ID")
	}

	userID, _ := utils.GetUserIDFromLocals(c) // Views can be anonymous

	view := ShoutView{
		UserID:  userID,
		ShoutID: shoutID,
	}

	if err := view.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.is.createView(&view); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

/********
 * SAVE *
 ********/

func (h *handler) createSave(c *fiber.Ctx) error {
	shoutID, err := uuid.Parse(c.Params("shout_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid shout ID")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	save := ShoutSave{
		UserID:  userID,
		ShoutID: shoutID,
	}

	if err := save.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.is.createSave(&save); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) deleteSave(c *fiber.Ctx) error {
	shoutID, err := uuid.Parse(c.Params("shout_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid shout ID")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	if err := h.is.deleteSave(shoutID, userID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

/****************
 * INTERACTIONS *
 ****************/

func (h *handler) getInteractions(c *fiber.Ctx) error {
	shoutID, err := uuid.Parse(c.Params("shout_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid shout ID")
	}

	result, err := h.is.getInteractions(shoutID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(result)
}
