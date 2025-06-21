package walkthrough

import "github.com/gofiber/fiber/v2"

func (h *handler) RegisterRoutes(router fiber.Router) {
	group := router.Group("/walkthroughs")

	group.Post("/", h.create)
	group.Get("/", h.list)
	group.Get("/:id", h.get)
	group.Put("/:id", h.update)
	group.Delete("/:id", h.delete)

	group.Post("/:id/steps", h.addStep)
	group.Put("/steps/:step_id", h.updateStep)
	group.Delete("/steps/:step_id", h.deleteStep)

	group.Post("/:id/comments", h.addComment)
	group.Get("/:id/comments", h.listComments)
}
