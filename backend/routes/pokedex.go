package routes

import (
	"pokemon/controllers"
	"pokemon/middleware"

	"github.com/gofiber/fiber/v2"
)

var ContentManagerRoles = []string{"admin", "moderator"}
func RegisterPokedexRoutes(api fiber.Router) {
	pokedex := api.Group("/me/pokedex")

	pokedex.Get("/:game_id", controllers.GetPokedexByGame)

	admin := pokedex.Group("/", middleware.RequireAuth(), middleware.RequireRole(ContentManagerRoles...))
	admin.Post("/:game_id/:pokemon_id", controllers.UpsertPokedexEntry)
	admin.Put("/entry/:entry_id", controllers.UpdatePokedexEntry)
	admin.Delete("/entry/:entry_id", controllers.DeletePokedexEntry)
}
