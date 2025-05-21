package routes

import (
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router) {
    auth := api.Group("/auth")
    auth.Post("/register", controllers.Register)
    auth.Post("/login", controllers.EmailLogin)
    auth.Post("/username-login", controllers.UsernameLogin)

    api.Get("/profile", middleware.RequireAuth(), controllers.Profile)
 
    admin := api.Group("/admin", middleware.RequireAuth(), middleware.RequireRole("admin", "moderator"))
	admin.Get("/users", controllers.GetAllUsers)
	admin.Patch("/users/:id/ban", controllers.BanUser)
}
