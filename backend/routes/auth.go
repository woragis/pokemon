package routes

import (
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app fiber.Router) {
    auth := app.Group("/auth")
    auth.Post("/register", controllers.Register)
    auth.Post("/login", controllers.EmailLogin)
    auth.Post("/username-login", controllers.UsernameLogin)

    app.Get("/profile", middleware.RequireAuth(), controllers.Profile)
}
