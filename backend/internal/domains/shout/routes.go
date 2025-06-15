package shout

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterRoutes(app fiber.Router) {
	shouts := app.Group("/shouts")
	shouts.Get("/", h.listShouts)
	shouts.Get("/:id", h.getShout)
	shouts.Get("/user/:userID", h.listShoutsByUser)
	shouts.Get("/parent/:reshoutID", h.listShoutsByParent)

	shouts.Use(middleware.AuthRequired())
	shouts.Post("/", h.createShout)
	shouts.Post("/:id", h.retweetShout)
	shouts.Put("/:id", h.updateShout)
	shouts.Delete("/:id", h.deleteShout)
}