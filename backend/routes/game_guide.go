package routes

import (
	"pokemon/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterGameGuideRoutes(app fiber.Router) {
	guides := app.Group("/guides")

	guides.Post("/", controllers.CreateGameGuide)
	guides.Get("/", controllers.ListGameGuides)
	guides.Get("/:slug", controllers.GetGameGuide)
	guides.Put("/:slug", controllers.UpdateGameGuide)

}