package game

import "github.com/gofiber/fiber/v2"

func (h *handler) RegisterRoutes(router fiber.Router) {
	group := router.Group("/games")

	group.Post("/", h.createGame)
	group.Get("/", h.listGames)
	group.Get("/:id", h.getGame)
	group.Put("/:id", h.updateGame)
	group.Delete("/:id", h.deleteGame)

	group.Get("/:id/pokedex", h.getUserGamePokedex)
	group.Post("/:id/pokedex", h.createUserGamePokedex)
	group.Put("/:id/pokedex", h.updateUserGamePokedex)
}
