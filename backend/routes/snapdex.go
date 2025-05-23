package routes

import (
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterSnapdexRoutes(api fiber.Router) {
	snaps := api.Group("/snapdex")

	snaps.Get("/", controllers.GetSnaps)
	snaps.Get("/me", middleware.RequireAuth(), controllers.GetOwnSnaps)
	snaps.Get("/:id", controllers.GetSnap)

	auth := snaps.Group("/", middleware.RequireAuth())
	auth.Post("/", controllers.CreateSnap)
	auth.Delete("/:id", controllers.DeleteSnap)
}
