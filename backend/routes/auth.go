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

    app.Get("/me", middleware.RequireAuth(), func(c *fiber.Ctx) error {
        userID := c.Locals("user_id")
        return c.JSON(fiber.Map{"user_id": userID})
    })
}
