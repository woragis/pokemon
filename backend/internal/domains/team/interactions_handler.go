package team

import (
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/**************************
 * HANDLER IMPLEMENTATION  *
 **************************/

// POST /teams/:id/save
func (h *handler) saveTeam(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	if err := h.i.saveTeam(userID, teamID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// DELETE /teams/:id/save
func (h *handler) unsaveTeam(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	if err := h.i.unsaveTeam(userID, teamID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GET /teams/saved
func (h *handler) getSavedTeams(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	limit := c.QueryInt("limit", 10)
	if limit < 1 {
		limit = 10
	}

	offset := c.QueryInt("offset", 0)
	if offset < 0 {
		offset = 0
	}

	teams, err := h.i.getSavedTeams(userID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(teams)
}

// GET /teams/:id/saved
func (h *handler) isTeamSavedByUser(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	saved, err := h.i.isTeamSavedByUser(userID, teamID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"saved": saved})
}

// POST /teams/:id/comments
func (h *handler) commentTeam(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	var body struct {
		Content  string    `json:"content"`
		ParentID *uuid.UUID `json:"parent_id,omitempty"`
	}
	if err := c.BodyParser(&body); err != nil || body.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	err = h.i.commentTeam(userID, teamID, body.Content, body.ParentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// GET /teams/:id/comments
func (h *handler) getTeamComments(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	limit := c.QueryInt("limit", 10)
	if limit < 1 {
		limit = 10
	}

	offset := c.QueryInt("offset", 0)
	if offset < 0 {
		offset = 0
	}

	comments, err := h.i.getTeamComments(teamID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(comments)
}

// GET /teams/:id/comments/count
func (h *handler) getTeamCommentCount(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	count, err := h.i.getTeamCommentCount(teamID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"count": count})
}

// PUT /comments/:id
func (h *handler) updateComment(c *fiber.Ctx) error {
	commentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid comment ID"})
	}

	var comment TeamComment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	// Make sure comment ID is set properly to avoid accidental overwrite
	comment.ID = commentID

	err = h.i.updateComment(&comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DELETE /comments/:id
func (h *handler) deleteComment(c *fiber.Ctx) error {
	commentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid comment ID"})
	}

	// Need teamID to invalidate cache — ideally pass it via query param or header
	teamIDStr := c.Query("team_id")
	if teamIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "team ID required"})
	}
	teamID, err := uuid.Parse(teamIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	err = h.i.deleteComment(commentID, teamID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// POST /teams/:id/like
func (h *handler) likeTeam(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	if err := h.i.likeTeam(userID, teamID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// DELETE /teams/:id/like
func (h *handler) unlikeTeam(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	if err := h.i.unlikeTeam(userID, teamID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GET /teams/:id/likes/count
func (h *handler) getTeamLikeCount(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	count, err := h.i.getTeamLikeCount(teamID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"likes": count})
}

// GET /teams/:id/likes
func (h *handler) isTeamLikedByUser(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	liked, err := h.i.isTeamLikedByUser(userID, teamID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"liked": liked})
}

// POST /teams/:id/view
func (h *handler) viewTeam(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	var userID *uuid.UUID
	if id, err := utils.GetUserIDFromLocals(c); err == nil {
		userID = &id
	}

	if err := h.i.viewTeam(userID, teamID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// GET /teams/:id/views/count
func (h *handler) getTeamViewCount(c *fiber.Ctx) error {
	teamID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	count, err := h.i.getTeamViewCount(teamID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"views": count})
}
