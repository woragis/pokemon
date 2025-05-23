package routes

import (
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)
func RegisterForumRoutes(api fiber.Router) {
	forum := api.Group("/forum")

	// Topics
	forum.Post("/topics", middleware.RequireAuth(), controllers.CreateForumTopic)
	forum.Get("/topics", controllers.GetForumTopics)
	forum.Get("/topics/:id", controllers.GetForumTopicById)
	forum.Put("/topics/:id", middleware.RequireAuth(), controllers.EditForumTopicById)
	forum.Delete("/topics/:id", middleware.RequireAuth(), controllers.DeleteForumTopicById)
	// Topic interaction
	forum.Post("/topics/:id/like", middleware.RequireAuth(), controllers.LikeForumTopic)
	forum.Post("/topics/:id/like", middleware.RequireAuth(), controllers.UnlikeForumTopic)
	forum.Post("/topics/:id/view", middleware.RequireAuth(), controllers.ViewForumTopic)
	forum.Get("/topics/:id/comments", controllers.GetForumTopicComments)
	forum.Post("/topics/:id/comments", middleware.RequireAuth(), controllers.CommentOnForumTopic)

	// Categories
	forum.Post("/categories", middleware.RequireAuth(), controllers.CreateForumCategory)
	forum.Get("/categories", controllers.GetForumCategories)
	forum.Get("/categories/:id", controllers.GetForumCategoryById)
	forum.Put("/categories/:id", middleware.RequireAuth(), controllers.EditForumCategoryById)
	forum.Delete("/categories/:id", middleware.RequireAuth(), controllers.DeleteForumCategoryById)
}
