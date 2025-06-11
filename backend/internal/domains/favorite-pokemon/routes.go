package favoritepokemon

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterRoutes(app fiber.Router) {
	favMon := app.Group("/favorite-pokemons")

	// favMon.Get("/", h.listRecent)
	favMon.Get("/popular", h.listByPopular)
	favMon.Get("/user/:user_id", h.listByUser)
	favMon.Get("/:id", h.getByID)

	favMon.Use(middleware.AuthRequired())
	favMon.Post("/", h.create)
	favMon.Put("/:id", h.update)
	favMon.Delete("/:id", h.delete)
}