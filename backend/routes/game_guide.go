package routes

import (
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterGameGuideRoutes(api fiber.Router) {
	guides := api.Group("/guides")

	guides.Get("/", controllers.ListGameGuides)
	guides.Get("/:slug", controllers.GetGameGuide)

	admin := guides.Group("/", middleware.RequireAuth(), middleware.RequireRole(ContentManagerRoles...))
	admin.Post("/", controllers.CreateGameGuide)
	admin.Put("/:slug", controllers.UpdateGameGuide)
}
