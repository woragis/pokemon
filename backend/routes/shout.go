package routes

import (
	"pokemon/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterShoutRoutes(app fiber.Router) {
	shouts := app.Group("/shouts")

	// Create shout
	shouts.Post("/", controllers.PostShout)

	// Feed options
	shouts.Get("/", controllers.GetShoutFeed)              // Basic feed
	shouts.Get("/page", controllers.GetPaginatedFeed)      // Page-based feed
	shouts.Get("/infinite", controllers.GetInfiniteFeed)   // Cursor-based feed (infinite scroll)

	// Interactions
	shouts.Post("/:id/like", controllers.LikeShout)
	shouts.Post("/:id/comment", controllers.CommentOnShout)
	shouts.Post("/:id/reshout", controllers.Reshout)

	// Timeline
	shouts.Get("/user/:id", controllers.GetUserTimeline)

	// Delete reshout
	shouts.Delete("/:id/reshout", controllers.DeleteReshout)

	// (Optional) AI reply
	// shouts.Post("/:id/ai-reply", controllers.AIReplyToShout)
}
