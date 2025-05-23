package routes

import (
	"pokemon/config"
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupPermissionRoutes(api fiber.Router) {
	permission := api.Group("/permissions")

	permission.Get("/", controllers.GetPermissions)
	permission.Get("/:id", controllers.GetPermission)

	admin := permission.Group("/admin/permissions", middleware.RequireAuth(), middleware.RequireRole(config.Admins...))
	admin.Post("/", controllers.CreatePermission)
	admin.Put("/:id", controllers.UpdatePermission)
	admin.Delete("/:id", controllers.DeletePermission)
}