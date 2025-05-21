package routes

import (
	"pokemon/config"
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func BlogRoutes(api fiber.Router) {
	blog := api.Group("/blog")

	blog.Get("/", controllers.GetBlogPosts)         // GET all
	blog.Get("/:id", controllers.GetBlogPost)       // GET by ID

	writers := blog.Group("/", middleware.RequireAuth(), middleware.RequireRole(config.Writers...))
	writers.Post("/", controllers.CreateBlogPost)      // POST new
	writers.Put("/:id", controllers.UpdateBlogPost)    // PUT update
	writers.Delete("/:id", controllers.DeleteBlogPost) // DELETE
}
