package team

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterTeamRoutes(app *fiber.App, handler TeamHandler) {
	teamGroup := app.Group("/teams")

	teamGroup.Get("/:id", handler.GetTeam)
	teamGroup.Get("/user/:user_id", handler.ListTeams)

	teamGroup.Use(middleware.AuthRequired())
	teamGroup.Post("/", handler.CreateTeam)
	teamGroup.Put("/:id", handler.UpdateTeam)
	teamGroup.Delete("/:id", handler.DeleteTeam)
}
