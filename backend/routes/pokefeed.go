package routes

import (
	"pokemon/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPokeFeedRoutes(app *fiber.App) {
    posts := app.Group("/pokefeed")

    posts.Post("/", controllers.PostPokePost)            // Create post
    posts.Get("/", controllers.GetPokeFeed)              // Get feed (paginated)
    posts.Get("/user/:id", controllers.GetUserPokePosts) // Get user posts
    posts.Post("/:id/like", controllers.LikePokePost)    // Like post
}
