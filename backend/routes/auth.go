package routes

import (
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupAuthRoutes(app *fiber.App, db *gorm.DB) {
    auth := app.Group("/auth")
    auth.Post("/register", controllers.Register)
    auth.Post("/login", controllers.Login)

    app.Get("/me", middleware.Protected(), func(c *fiber.Ctx) error {
        userID := c.Locals("user_id")
        return c.JSON(fiber.Map{"user_id": userID})
    })
}
