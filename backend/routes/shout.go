package routes

import (
	"pokemon/config"
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterShoutRoutes(api fiber.Router) {
	shouts := api.Group("/shouts")

	// Create shout
	shouts.Post("/", middleware.RequireAuth(), controllers.PostShout)

	// Feed options
	shouts.Get("/", controllers.GetShoutFeed)              // Basic feed
	shouts.Get("/page", controllers.GetPaginatedFeed)      // Page-based feed
	shouts.Get("/infinite", controllers.GetInfiniteFeed)   // Cursor-based feed (infinite scroll)
	
	// Timeline
	shouts.Get("/user/:id", controllers.GetUserTimeline)

	// Interactions
	userShouts := shouts.Group("/", middleware.RequireAuth())
	userShouts.Post("/:id/like", controllers.LikeShout)
	userShouts.Post("/:id/comment", controllers.CommentOnShout)
	userShouts.Post("/:id/reshout", controllers.Reshout)

	// Delete reshout
	userShouts.Delete("/:id/reshout", controllers.DeleteReshout)

	// (Optional) AI reply
	// shouts.Post("/:id/ai-reply", controllers.AIReplyToShout)
	admin := api.Group("/admin", middleware.RequireRole(), middleware.RequireRole(config.Moderators...))

	// New shout moderation routes
	admin.Get("/shouts", controllers.ModerateShouts)
	admin.Delete("/shouts/:id", controllers.DeleteShout)
}
