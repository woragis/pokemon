package routes

import (
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterPokeFeedRoutes(api fiber.Router) {
    posts := api.Group("/pokefeed")

    posts.Get("/", controllers.GetPokeFeed)              // Get feed (paginated)
    posts.Get("/user/:id", controllers.GetUserPokePosts) // Get user posts

    userPosts := posts.Group("/", middleware.RequireAuth())
    userPosts.Post("/", controllers.PostPokePost)            // Create post
    userPosts.Post("/:id/like", controllers.LikePokePost)    // Like post
    userPosts.Post("/:id/comment", controllers.CommentOnPokePost)
    userPosts.Post("/:id/unlike", controllers.UnlikePokePost)

    // Review to see if they would be in this scope
    users := api.Group("/trainers", middleware.RequireAuth())

    users.Post("/:id/follow", controllers.FollowTrainer)
    users.Post("/:id/unfollow", controllers.UnfollowTrainer)
}
