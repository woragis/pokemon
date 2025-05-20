package routes

import (
	"pokemon/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterPokedexRoutes(router fiber.Router) {
	pokedex := router.Group("/me/pokedex")

	pokedex.Get("/:game_id", controllers.GetPokedexByGame)
	pokedex.Post("/:game_id/:pokemon_id", controllers.UpsertPokedexEntry)
	pokedex.Put("/entry/:entry_id", controllers.UpdatePokedexEntry)
	pokedex.Delete("/entry/:entry_id", controllers.DeletePokedexEntry)
}
