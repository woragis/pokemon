package guide

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *gameGuideHandler) RegisterRoutes(app fiber.Router) {
	guides := app.Group("/game-guides")

	// Public routes
	guides.Get("/", h.list)
	guides.Get("/:id", h.getByID)
	guides.Get("/slug/:slug", h.getBySlug)
	guides.Get("/user/:user_id", h.listByUser)

	// Authenticated routes
	guides.Use(middleware.AuthRequired())
	guides.Post("/", h.create)
	guides.Put("/:id", h.update)
	guides.Delete("/:id", h.delete)
}
