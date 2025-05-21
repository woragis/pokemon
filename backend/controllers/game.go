package controllers

import (
	"pokemon/database"
	"pokemon/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GET /pokemon-games
func GetPokemonGames(c *fiber.Ctx) error {
	var games []models.PokemonGame
	if err := database.DB.Find(&games).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching games",
		})
	}
	return c.JSON(games)
}

// GET /pokemon-games/:id
func GetPokemonGame(c *fiber.Ctx) error {
	id := c.Params("id")
	var game models.PokemonGame

	if err := database.DB.First(&game, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Game not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching game",
		})
	}

	return c.JSON(game)
}

// POST /pokemon-games
func CreatePokemonGame(c *fiber.Ctx) error {
	var input models.PokemonGame
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	input.ID = uuid.New()

	if err := database.DB.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating game",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(input)
}

// PUT /pokemon-games/:id
func UpdatePokemonGame(c *fiber.Ctx) error {
	id := c.Params("id")
	var game models.PokemonGame

	if err := database.DB.First(&game, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Game not found",
		})
	}

	var input models.PokemonGame
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// update fields
	game.Name = input.Name
	game.Generation = input.Generation
	game.ReleasedAt = input.ReleasedAt
	game.Description = input.Description

	if err := database.DB.Save(&game).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating game",
		})
	}

	return c.JSON(game)
}

// DELETE /pokemon-games/:id
func DeletePokemonGame(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := database.DB.Delete(&models.PokemonGame{}, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting game",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
