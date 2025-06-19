package guide

import (
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type gameGuideHandler struct {
	s gameGuideService
}

func NewHandler(db *gorm.DB, redis *redis.Client) *gameGuideHandler {
	repo := newGameGuideRepo(db)
	serv := newGameGuideService(repo, redis)

	return &gameGuideHandler{s: serv}
}

func (h *gameGuideHandler) create(c *fiber.Ctx) error {
	var guide GameGuide
	if err := c.BodyParser(&guide); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	guide.AuthorID = userID
	if err := guide.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.s.create(&guide); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(guide)
}

func (h *gameGuideHandler) getByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	guide, err := h.s.getByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(guide)
}

func (h *gameGuideHandler) getBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing slug")
	}

	guide, err := h.s.getBySlug(slug)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(guide)
}

func (h *gameGuideHandler) update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	var guide GameGuide
	if err := c.BodyParser(&guide); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	guide.ID = id
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}
	guide.AuthorID = userID

	if err := guide.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.s.update(&guide); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(guide)
}

func (h *gameGuideHandler) delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	if err := h.s.delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *gameGuideHandler) list(c *fiber.Ctx) error {
	limit, offset := utils.ParsePagination(c)
	guides, err := h.s.list(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(guides)
}

func (h *gameGuideHandler) listByUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id")
	}

	limit, offset := utils.ParsePagination(c)
	guides, err := h.s.listByAuthor(userID, limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(guides)
}
