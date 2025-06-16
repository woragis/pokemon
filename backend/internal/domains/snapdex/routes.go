package snapdex

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterRoutes(app fiber.Router) {
	snaps := app.Group("/snaps")

	snaps.Get("/", h.listSnaps)
	snaps.Get("/:id", h.getSnap)
	snaps.Get("/user/:user_id", h.listSnapsByUser)
	snaps.Get("/user/:user_id/count", h.countSnapsByUser)

	snaps.Use(middleware.AuthRequired())
	snaps.Post("/", h.createSnap)
	snaps.Put("/:id", h.updateSnap)
	snaps.Delete("/:id", h.deleteSnap)
}
