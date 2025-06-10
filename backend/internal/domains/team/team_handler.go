package team

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*********************
 * HANDLER INTERFACE *
 *********************/

type TeamHandler interface {
	CreateTeam(c *fiber.Ctx) error
	GetTeam(c *fiber.Ctx) error
	ListTeams(c *fiber.Ctx) error
	UpdateTeam(c *fiber.Ctx) error
	DeleteTeam(c *fiber.Ctx) error
}

/**************************
 * HANDLER IMPLEMENTATION *
 **************************/

 type teamHandler struct {
	service TeamService
}

func NewTeamHandler(service TeamService) TeamHandler {
	return &teamHandler{service}
}

// POST /teams
func (h *teamHandler) CreateTeam(c *fiber.Ctx) error {
	var team Team
	if err := c.BodyParser(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if err := h.service.CreateTeam(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(team)
}

// GET /teams/:id
func (h *teamHandler) GetTeam(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}
	team, err := h.service.GetTeam(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "team not found"})
	}
	return c.JSON(team)
}

// GET /teams/user/:user_id
func (h *teamHandler) ListTeams(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}
	teams, err := h.service.ListTeams(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(teams)
}

// PUT /teams/:id
func (h *teamHandler) UpdateTeam(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	var team Team
	if err := c.BodyParser(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	team.ID = id

	if err := h.service.UpdateTeam(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(team)
}

// DELETE /teams/:id
func (h *teamHandler) DeleteTeam(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}
	if err := h.service.DeleteTeam(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}