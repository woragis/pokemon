package favoritepokemon

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type handler struct {
	s favoritepokemonService
}

func NewHandler(db *gorm.DB, redis *redis.Client) *handler {
	repo := newFavoritePokemonRepository(db)
	service := newFavoritePokemonService(repo, redis)
	return &handler{s: service}
}

func (h *handler) create(c *fiber.Ctx) error {
	var mon FavoritePokemon
	if err := c.BodyParser(&mon); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON")
	}

	mon.ID = uuid.New() // ensure it's new
	if err := h.s.create(&mon); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(mon)
}

func (h *handler) listByPopular(c *fiber.Ctx) error {
	mons, err := h.s.listByPopular()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(mons)
}

func (h *handler) listByUser(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid user ID")
	}

	mons, err := h.s.listByUser(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(mons)
}

func (h *handler) getByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID format")
	}

	mon, err := h.s.getByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(mon)
}

func (h *handler) update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	var mon FavoritePokemon
	if err := c.BodyParser(&mon); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON")
	}

	mon.ID = id // enforce path ID
	if err := h.s.update(&mon); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(mon)
}

func (h *handler) delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid ID")
	}

	if err := h.s.delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

