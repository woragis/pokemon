package routes

import (
	"pokemon/config"
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterPokedexRoutes(api fiber.Router) {
	pokedex := api.Group("/me/pokedex")

	pokedex.Get("/:game_id", controllers.GetPokedexByGame)

	admin := pokedex.Group("/", middleware.RequireAuth(), middleware.RequireRole(config.Moderators...))
	admin.Post("/:game_id/:pokemon_id", controllers.UpsertPokedexEntry)
	admin.Put("/entry/:entry_id", controllers.UpdatePokedexEntry)
	admin.Delete("/entry/:entry_id", controllers.DeletePokedexEntry)
}
