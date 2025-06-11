package team

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

/**************************
 * HANDLER IMPLEMENTATION *
 **************************/

 type handler struct {
	service teamService
}

func NewHandler(db *gorm.DB, redis *redis.Client) *handler {
	repo := newRepository(db)
	service := newService(repo, redis)
	return &handler{service}
}

// POST /teams
func (h *handler) createTeam(c *fiber.Ctx) error {
	var team Team
	if err := c.BodyParser(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if err := h.service.createTeam(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(team)
}

// GET /teams/:id
func (h *handler) getTeam(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}
	team, err := h.service.getTeam(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "team not found"})
	}
	return c.JSON(team)
}

// GET /teams/user/:user_id
func (h *handler) listTeams(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("user_id"))
	limit := 5
	offset := 0
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}
	teams, err := h.service.listTeams(userID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(teams)
}

// PUT /teams/:id
func (h *handler) updateTeam(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	var team Team
	if err := c.BodyParser(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	team.ID = id

	if err := h.service.updateTeam(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(team)
}

// DELETE /teams/:id
func (h *handler) deleteTeam(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}
	if err := h.service.deleteTeam(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}