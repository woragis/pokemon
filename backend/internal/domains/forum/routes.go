package forum

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterRoutes(app fiber.Router) {
	forumGroup := app.Group("/forum/topics")
	forumGroup.Get("/", h.list)
	forumGroup.Get("/user/:user_id", h.listByUser)
	forumGroup.Get("/:id", h.getByID)

	forumGroup.Use(middleware.AuthRequired())

	forumGroup.Post("/", h.create)
	forumGroup.Put("/:id", h.update)
	forumGroup.Delete("/:id", h.delete)
}
