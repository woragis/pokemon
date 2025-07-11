package news

import (
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type handler struct {
	s   service
	// viewCache viewService
}

func NewHandler(db *gorm.DB, redis *redis.Client) *handler {
	repo := newRepo(db)
	serv := newServ(repo, redis)
	return &handler{s: serv}
}

// POST /news
func (h *handler) create(c *fiber.Ctx) error {
	var news News

	if err := c.BodyParser(&news); err != nil {
		return fiber.ErrBadRequest
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}
	news.UserID = userID

	if err := news.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.s.create(&news); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(news)
}

// GET /news/:id
func (h *handler) get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	news, err := h.s.get(id)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(news)
}

// GET /news
func (h *handler) list(c *fiber.Ctx) error {
	limit, offset := utils.ParsePagination(c)

	news, err := h.s.list(limit, offset)
	if err != nil {
		return err
	}
	return c.JSON(news)
}

// GET /news/user/:user_id
func (h *handler) listByUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return fiber.ErrBadRequest
	}
	limit, offset := utils.ParsePagination(c)

	news, err := h.s.listByUser(userID, limit, offset)
	if err != nil {
		return err
	}
	return c.JSON(news)
}

// PUT /news/:id
func (h *handler) update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var news News
	if err := c.BodyParser(&news); err != nil {
		return fiber.ErrBadRequest
	}

	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	news.ID = id
	news.UserID = userID

	if err := news.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.s.update(&news); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// DELETE /news/:id
func (h *handler) delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.s.delete(id); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// // POST /news/:id/view
// func (h *handler) createView(c *fiber.Ctx) error {
// 	newsID, err := uuid.Parse(c.Params("id"))
// 	if err != nil {
// 		return fiber.ErrBadRequest
// 	}

// 	view := &NewsView{NewsID: newsID}
// 	userID, _ := utils.GetUserIDFromLocals(c) // optional
// 	view.UserID = userID

// 	if err := view.Validate(); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	if err := h.viewCache.create(view); err != nil {
// 		return err
// 	}
// 	return c.SendStatus(fiber.StatusCreated)
// }
