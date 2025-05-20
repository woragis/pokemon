package routes

import (
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin", middleware.RequireRole("admin", "moderator"))

	admin.Get("/users", controllers.GetAllUsers)
	admin.Patch("/users/:id/ban", controllers.BanUser)

	// New shout moderation routes
	admin.Get("/shouts", controllers.ModerateShouts)
	admin.Delete("/shouts/:id", controllers.DeleteShout)
}
