package game

import (
	"pokemon/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type handler struct {
	gameSvc    gameService
	pokedexSvc gamePokedexService
}

func NewHandler(db *gorm.DB, redis *redis.Client) *handler {
	gameR := newGameRepo(db)
	gameS := newGameService(gameR, redis)

	pokedexR := newGamePokedexRepo(db)
	pokedexS := newGamePokedexService(pokedexR, redis)

	return &handler{
		gameSvc: gameS,
		pokedexSvc: pokedexS,
	}
}

func (h *handler) createGame(c *fiber.Ctx) error {
	var game Game
	if err := c.BodyParser(&game); err != nil {
		return fiber.ErrBadRequest
	}
	if err := h.gameSvc.create(c.Context(), &game); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(game)
}

func (h *handler) listGames(c *fiber.Ctx) error {
	limit, offset := utils.ParsePagination(c)
	list, total, err := h.gameSvc.list(c.Context(), limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"total": total,
		"items": list,
	})
}

func (h *handler) getGame(c *fiber.Ctx) error {
	id := c.Params("id")
	game, err := h.gameSvc.getByID(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "game not found")
	}
	return c.JSON(game)
}

func (h *handler) updateGame(c *fiber.Ctx) error {
	var game Game
	id := c.Params("id")
	if err := c.BodyParser(&game); err != nil {
		return fiber.ErrBadRequest
	}
	game.ID = utils.ParseUUID(id)
	if err := h.gameSvc.update(c.Context(), &game); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(game)
}

func (h *handler) deleteGame(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.gameSvc.delete(c.Context(), id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *handler) getUserGamePokedex(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}
	gameID := c.Params("id")

	dex, err := h.pokedexSvc.getByUserAndGame(c.Context(), userID.String(), gameID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "pokedex not found")
	}
	return c.JSON(dex)
}

func (h *handler) createUserGamePokedex(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}
	gameID := c.Params("id")

	var dex GamePokedex
	if err := c.BodyParser(&dex); err != nil {
		return fiber.ErrBadRequest
	}

	dex.GameID = utils.ParseUUID(gameID)
	dex.UserID = userID

	if err := h.pokedexSvc.create(c.Context(), &dex); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(dex)
}

func (h *handler) updateUserGamePokedex(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromLocals(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}
	gameID := c.Params("id")

	var dex GamePokedex
	if err := c.BodyParser(&dex); err != nil {
		return fiber.ErrBadRequest
	}

	dex.GameID = utils.ParseUUID(gameID)
	dex.UserID = userID

	if err := h.pokedexSvc.update(c.Context(), &dex); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(dex)
}
