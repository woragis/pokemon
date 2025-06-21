package walkthrough

import (
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type handler struct {
	s service
}

func NewHandler(db *gorm.DB, redis *redis.Client) *handler {
	repo := newRepo(db)
	serv := newServ(repo, redis)

	return &handler{s: serv}
}

func (h *handler) create(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	var payload Walkthrough
	if err := c.BodyParser(&payload); err != nil {
		return fiber.ErrBadRequest
	}

	payload.UserID = userID

	if err := h.s.createWalkthrough(c.Context(), &payload); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(payload)
}

func (h *handler) get(c *fiber.Ctx) error {
	id := c.Params("id")
	wt, err := h.s.getWalkthrough(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Walkthrough not found")
	}
	return c.JSON(wt)
}

func (h *handler) list(c *fiber.Ctx) error {
	limit, offset := utils.ParsePagination(c)
	list, total, err := h.s.listWalkthroughs(c.Context(), limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"total": total,
		"items": list,
	})
}

func (h *handler) update(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	id := c.Params("id")
	wt, err := h.s.getWalkthrough(c.Context(), id)
	if err != nil {
		return fiber.ErrNotFound
	}

	if wt.UserID != userID {
		return fiber.ErrForbidden
	}

	if err := c.BodyParser(&wt); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.s.updateWalkthrough(c.Context(), wt); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(wt)
}

func (h *handler) delete(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	id := c.Params("id")
	wt, err := h.s.getWalkthrough(c.Context(), id)
	if err != nil {
		return fiber.ErrNotFound
	}

	if wt.UserID != userID {
		return fiber.ErrForbidden
	}

	if err := h.s.deleteWalkthrough(c.Context(), id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *handler) addStep(c *fiber.Ctx) error {
	var step WalkthroughStep
	if err := c.BodyParser(&step); err != nil {
		return fiber.ErrBadRequest
	}

	step.WalkthroughID = uuid.MustParse(c.Params("id"))

	if err := h.s.addStep(c.Context(), &step); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(step)
}

func (h *handler) updateStep(c *fiber.Ctx) error {
	var step WalkthroughStep
	if err := c.BodyParser(&step); err != nil {
		return fiber.ErrBadRequest
	}
	step.ID = uuid.MustParse(c.Params("step_id"))

	if err := h.s.updateStep(c.Context(), &step); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(step)
}

func (h *handler) deleteStep(c *fiber.Ctx) error {
	stepID := c.Params("step_id")
	if err := h.s.deleteStep(c.Context(), stepID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *handler) addComment(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	var comment WalkthroughComment
	if err := c.BodyParser(&comment); err != nil {
		return fiber.ErrBadRequest
	}

	comment.WalkthroughID = uuid.MustParse(c.Params("id"))
	comment.UserID = userID

	if err := h.s.addComment(c.Context(), &comment); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(comment)
}

func (h *handler) listComments(c *fiber.Ctx) error {
	id := c.Params("id")
	limit, offset := utils.ParsePagination(c)
	comments, total, err := h.s.listComments(c.Context(), id, limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"total": total,
		"items": comments,
	})
}
