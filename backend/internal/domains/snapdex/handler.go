package snapdex

import (
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type handler struct {
	s SnapService
}

func NewHandler(db *gorm.DB, redis *redis.Client) handler {
	repo := newRepo(db)
	serv := newSnapService(repo, redis)

	return handler{s: serv}
}

// Private: Create Snap
func (h *handler) createSnap(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	var snap Snap
	if err := c.BodyParser(&snap); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	snap.UserID = userID
	if err := snap.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.s.create(&snap); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(snap)
}

// Private: Update Snap
func (h *handler) updateSnap(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid snap ID")
	}

	var snap Snap
	if err := c.BodyParser(&snap); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	snap.ID = id
	snap.UserID = userID

	if err := snap.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.s.update(&snap); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(snap)
}

// Private: Delete Snap
func (h *handler) deleteSnap(c *fiber.Ctx) error {
	_, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid snap ID")
	}

	if err := h.s.delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Public: Get Snap by ID
func (h *handler) getSnap(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid snap ID")
	}

	snap, err := h.s.getByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Snap not found")
	}

	return c.JSON(snap)
}

// Public: List All Snaps (paginated)
func (h *handler) listSnaps(c *fiber.Ctx) error {
	limit, offset := utils.ParsePagination(c)
	snaps, err := h.s.list(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(snaps)
}

// Public: List Snaps by User
func (h *handler) listSnapsByUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	limit, offset := utils.ParsePagination(c)
	snaps, err := h.s.listByUser(userID, limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(snaps)
}

// Public: Count Snaps by User
func (h *handler) countSnapsByUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	count, err := h.s.countByUser(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"count": count})
}
