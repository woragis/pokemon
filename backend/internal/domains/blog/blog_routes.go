package blog

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterRoutes(app fiber.Router) {
	blogGroup := app.Group("/blog")

	// Public Routes
	blogGroup.Get("/", h.listPosts)
	blogGroup.Get("/user/user_id", h.listPostsByUser)
	blogGroup.Get("/:id", h.getPost)
	// blogGroup.Get("/:id/comments", h.getComments)
	// blogGroup.Get("/:id/comments/count", h.getCommentsCount)
	// blogGroup.Get("/:id/views/count", h.getBlogViewCount)
	// blogGroup.Get("/:id/likes/count", h.getBlogLikeCount)

	// Auth Required
	blogGroup.Use(middleware.AuthRequired())

	// Blog CRUD
	blogGroup.Post("/", h.createPost)
	blogGroup.Put("/:id", h.updatePost)
	blogGroup.Delete("/:id", h.deletePost)
}