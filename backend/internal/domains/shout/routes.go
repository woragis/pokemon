package shout

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterRoutes(app fiber.Router) {
	shouts := app.Group("/shouts")

	// Public routes
	shouts.Get("/", h.listShouts)
	shouts.Get("/:id", h.getShout)
	shouts.Get("/user/:userID", h.listShoutsByUser)
	shouts.Get("/parent/:reshoutID", h.listShoutsByParent)

	// Interaction routes (public)
	shouts.Get("/:shout_id/interactions", h.getInteractions)
	shouts.Post("/:shout_id/views", middleware.AuthOptional(), h.createView)

	// Authenticated routes
	shouts.Use(middleware.AuthRequired())

	shouts.Post("/", h.createShout)
	shouts.Post("/:id/reshout", h.retweetShout)
	shouts.Put("/:id", h.updateShout)
	shouts.Delete("/:id", h.deleteShout)

	// Likes
	shouts.Post("/:shout_id/likes", h.createLike)
	shouts.Delete("/:shout_id/likes", h.deleteLike)

	// Comments
	shouts.Post("/comments", h.createComment)
	shouts.Put("/comments", h.updateComment)
	shouts.Delete("/comments/:comment_id", h.deleteComment)

	// Saves
	shouts.Post("/:shout_id/saves", h.createSave)
	shouts.Delete("/:shout_id/saves", h.deleteSave)
}
