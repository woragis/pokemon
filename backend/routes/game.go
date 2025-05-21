package routes

import (
	"pokemon/config"
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupGameRoutes(api fiber.Router) {
	games := api.Group("/games")
	games.Get("/", controllers.GetPokemonGames)
	games.Get("/:id", controllers.GetPokemonGame)

	mods := games.Group("/", middleware.RequireAuth(), middleware.RequireRole(config.Moderators...))
	mods.Post("/", controllers.CreatePokemonGame)
	mods.Put("/:id", controllers.UpdatePokemonGame)
	mods.Delete("/:id", controllers.DeletePokemonGame)
}