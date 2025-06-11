package team

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) RegisterRoutes(app fiber.Router) {
	teamGroup := app.Group("/teams")

	teamGroup.Get("/:id", handler.getTeam)
	teamGroup.Get("/user/:user_id", handler.listTeams)

	teamGroup.Use(middleware.AuthRequired())
	teamGroup.Post("/", handler.createTeam)
	teamGroup.Put("/:id", handler.updateTeam)
	teamGroup.Delete("/:id", handler.deleteTeam)
}
