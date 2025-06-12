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

	// commentGroup := forumGroup.Group("/comments")
	// commentGroup.Get("/", h)
	// commentGroup.Get("/:id", h.getCategoryByID)

	// commentGroup.Use(middleware.AuthRequired())

	// commentGroup.Post("/", h.createCategory)
	// commentGroup.Put("/:id", h.updateCategory)
	// commentGroup.Delete("/:id", h.deleteCategory)

	commentLikeGroup := forumGroup.Group("/comments/:comment_id/likes")
	commentLikeGroup.Get("/", h.getLikes)                    // get all likes for comment (or user-specific if JWT)
	commentLikeGroup.Get("/count", h.countLikes)             // total likes/dislikes
	commentLikeGroup.Use(middleware.AuthRequired())
	commentLikeGroup.Post("/", h.createLike)                 // body has {like: true/false}
	commentLikeGroup.Put("/", h.updateLike)                  // same
	commentLikeGroup.Delete("/", h.deleteLike)               // remove user's like
}
