package routes

import (
	"pokemon/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterShoutRoutes(app *fiber.App) {
	shouts := app.Group("/shouts")

	shouts.Post("/", controllers.PostShout)
	shouts.Get("/", controllers.GetShoutFeed)
	shouts.Post("/:id/like", controllers.LikeShout)
	// Later: shouts.Post("/:id/comment", ...)
}
