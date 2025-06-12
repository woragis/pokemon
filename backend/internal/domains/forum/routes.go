package forum

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterRoutes(app fiber.Router) {
	forumGroup := app.Group("/forum")

	topicGroup := forumGroup.Group("/topics")
	topicGroup.Get("/", h.listTopic)
	topicGroup.Get("/user/:user_id", h.listTopicByUser)
	topicGroup.Get("/:id", h.getTopicByID)

	topicGroup.Use(middleware.AuthRequired())

	topicGroup.Post("/", h.createTopic)
	topicGroup.Put("/:id", h.updateTopic)
	topicGroup.Delete("/:id", h.deleteTopic)

	categoryGroup := forumGroup.Group("/categories")
	categoryGroup.Get("/", h.listCategories)
	categoryGroup.Get("/:id", h.getCategoryByID)

	categoryGroup.Use(middleware.AuthRequired())

	categoryGroup.Post("/", h.createCategory)
	categoryGroup.Put("/:id", h.updateCategory)
	categoryGroup.Delete("/:id", h.deleteCategory)
}
