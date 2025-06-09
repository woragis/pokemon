package user

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) RegisterRoutes(router fiber.Router) {
    users := router.Group("/users")
    
    // Public routes
    users.Post("/", h.CreateUser)
    users.Post("/login", h.Login)
    
    // Protected routes
    users.Use(middleware.AuthRequired())
    users.Get("/", h.ListUsers)
    users.Get("/:id", h.GetUser)
    users.Put("/:id", h.UpdateUser)
    users.Delete("/:id", h.DeleteUser)
}