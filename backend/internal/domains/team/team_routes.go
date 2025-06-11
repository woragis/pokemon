package team

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterRoutes(app fiber.Router) {
	teamGroup := app.Group("/teams")

	// Public routes
	teamGroup.Get("/:id", h.getTeam)
	teamGroup.Get("/user/:user_id", h.listTeams)
	teamGroup.Get("/:id/comments", h.getTeamComments)
	teamGroup.Get("/:id/comments/count", h.getTeamCommentCount)
	teamGroup.Get("/:id/views/count", h.getTeamViewCount)
	teamGroup.Get("/:id/likes/count", h.getTeamLikeCount)

	// Auth required
	teamGroup.Use(middleware.AuthRequired())

	// Team CRUD
	teamGroup.Post("/", h.createTeam)
	teamGroup.Put("/:id", h.updateTeam)
	teamGroup.Delete("/:id", h.deleteTeam)

	// Interaction routes
	teamGroup.Post("/:id/comments", h.commentTeam)
	teamGroup.Put("/comments/:comment_id", h.updateComment)
	teamGroup.Delete("/comments/:comment_id", h.deleteComment)

	teamGroup.Post("/:id/like", h.likeTeam)
	teamGroup.Delete("/:id/like", h.unlikeTeam)
	teamGroup.Get("/:id/likes", h.isTeamLikedByUser)

	teamGroup.Post("/:id/save", h.saveTeam)
	teamGroup.Delete("/:id/save", h.unsaveTeam)
	teamGroup.Get("/saved", h.getSavedTeams)
	teamGroup.Get("/:id/saved", h.isTeamSavedByUser)

	teamGroup.Post("/:id/view", h.viewTeam)
}
