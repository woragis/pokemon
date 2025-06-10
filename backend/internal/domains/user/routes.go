package user

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterRoutes(router fiber.Router) {
    users := router.Group("/users")
    
    // Public routes
    users.Post("/", h.createUser)
    users.Post("/login", h.login)
    
    // Protected routes
    users.Use(middleware.AuthRequired())
    users.Get("/", h.listUsers)
    users.Get("/:id", h.getUser)
    users.Put("/:id", h.updateUser)
    users.Delete("/:id", h.deleteUser)
}