package shout

import (
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// shoutHandler struct with shoutService injected
type handler struct {
	s shoutService
	is interactionService
}

func NewHandler(db *gorm.DB, redis *redis.Client) *handler {
	repo := newShoutRepo(db)      // assumes your repo constructor
	service := newService(repo, redis)

	iRepo := newInteractionRepo(db)
	iService := newInteractionService(iRepo, redis)
	return &handler{s: service, is: iService}
}

// POST /shouts
func (h *handler) createShout(c *fiber.Ctx) error {
	var shout Shout
	if err := c.BodyParser(&shout); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthenticated"})
	}

	// Override user ID from context, ignore any UserID in the body
	shout.UserID = userID

	if err := h.s.createShout(&shout); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(shout)
}

// GET /shouts/:id
func (h *handler) getShout(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid shout ID"})
	}
	shout, err := h.s.getShout(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "shout not found"})
	}
	return c.JSON(shout)
}

// GET /shouts
func (h *handler) listShouts(c *fiber.Ctx) error {
	limit, offset := c.QueryInt("limit", 10), c.QueryInt("offset", 0)

	shouts, err := h.s.listShouts(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(shouts)
}

// GET /shouts/user/:userID
func (h *handler) listShoutsByUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("userID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	limit, offset := c.QueryInt("limit", 10), c.QueryInt("offset", 0)

	shouts, err := h.s.listShoutsByUser(userID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(shouts)
}

// GET /shouts/parent/:reshoutID
func (h *handler) listShoutsByParent(c *fiber.Ctx) error {
	reshoutID, err := uuid.Parse(c.Params("reshoutID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid reshout ID"})
	}

	limit, offset := c.QueryInt("limit", 10), c.QueryInt("offset", 0)

	shouts, err := h.s.listShoutsByParent(reshoutID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(shouts)
}

// PUT /shouts/:id
func (h *handler) updateShout(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid shout ID"})
	}

	var shout Shout
	if err := c.BodyParser(&shout); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthenticated"})
	}

	// Override UserID and ensure the shout ID matches path
	shout.UserID = userID
	shout.ID = id

	if err := h.s.updateShout(&shout); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(shout)
}

// DELETE /shouts/:id
func (h *handler) deleteShout(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid shout ID"})
	}

	if err := h.s.deleteShout(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// POST /shouts/:id/reshout (retweet)
func (h *handler) retweetShout(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthenticated"})
	}

	originalID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid shout ID"})
	}

	newShout := &Shout{
		UserID:      userID,
		ReshoutOfID: &originalID,
		Content:     "", // retweet without comment
	}

	if err := h.s.createShout(newShout); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(newShout)
}
