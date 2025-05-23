package routes

import (
	"pokemon/config"
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router) {
	user := api.Group("/users")


	roles := user.Group("/roles")
	// user id
	roles.Get("/:id", controllers.GetUserRoles)
	roles.Post("/:id", controllers.AssignRoleToUser, middleware.RequireAllRoles(config.Admins...))
}