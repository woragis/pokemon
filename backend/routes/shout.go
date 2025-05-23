package routes

import (
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterShoutRoutes(api fiber.Router) {
	shouts := api.Group("/shouts")

	shouts.Get("/", controllers.GetShoutFeed)              // Basic feed
	shouts.Get("/page", controllers.GetPaginatedFeed)      // Page-based feed
	shouts.Get("/infinite", controllers.GetInfiniteFeed)   // Cursor-based feed (infinite scroll)
	shouts.Get("/user/:id", controllers.GetUserTimeline)
	shouts.Get("/:id", controllers.GetShoutByID)

	userShouts := shouts.Group("/", middleware.RequireAuth())
	userShouts.Post("/", middleware.RequireAuth(), controllers.PostShout)
	userShouts.Post("/:id/like", controllers.LikeShout)
	userShouts.Post("/:id/comment", controllers.CommentOnShout)
	userShouts.Post("/:id/reshout", controllers.Reshout)
	userShouts.Delete("/:id/reshout", controllers.DeleteReshout)

	// admin := shouts.Group("/admin", middleware.RequireRole(), middleware.RequireRole(config.Moderators...))
	// admin := api.Group("/shouts/admin")
	// admin.Post("/:id/flag")
	// admin.Get("/shouts", controllers.ModerateShouts)
	// admin.Delete("/shouts/:id", controllers.DeleteShout)
}
