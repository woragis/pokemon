package news

import "github.com/gofiber/fiber/v2"

func (h *handler) RegisterRoutes(app fiber.Router) {
	news := app.Group("/news")

	news.Get("/", h.list)
	news.Get("/:id", h.get)
	news.Get("/user/:user_id", h.listByUser)

	news.Post("/", h.create)
	news.Put("/:id", h.update)
	news.Delete("/:id", h.delete)

	// news.Post("/:id/view", h.createView)
}
