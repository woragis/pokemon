package routes

import (
	"pokemon/controllers"

	"github.com/gofiber/fiber/v2"
)

func BlogRoutes(app fiber.Router) {
	blog := app.Group("/blog")

	blog.Get("/", controllers.GetBlogPosts)         // GET all
	blog.Get("/:id", controllers.GetBlogPost)       // GET by ID
	blog.Post("/", controllers.CreateBlogPost)      // POST new
	blog.Put("/:id", controllers.UpdateBlogPost)    // PUT update
	blog.Delete("/:id", controllers.DeleteBlogPost) // DELETE
}
