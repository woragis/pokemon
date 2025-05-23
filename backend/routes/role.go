package routes

import (
	"pokemon/config"
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoleRoutes(api fiber.Router) {
	role := api.Group("/roles")

	role.Get("/", controllers.GetRoles)
	role.Get("/:id", controllers.GetRole)

	admin := role.Group("/", middleware.RequireAuth(), middleware.RequireRole(config.Admins...))
	admin.Post("/", controllers.CreateRole)
	admin.Put("/:id", controllers.UpdateRole)
	admin.Delete("/:id", controllers.DeleteRole)
}
