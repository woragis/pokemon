package controllers

import (
	"pokemon/database"
	"pokemon/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PokedexEntryInput struct {
	Caught         bool   `json:"caught"`
	Shiny          bool   `json:"shiny"`
	LivingDex      bool   `json:"living_dex"`
	ShinyLivingDex bool   `json:"shiny_living_dex"`
	Notes          string `json:"notes"`
}

// GET /me/pokedex/:game_id
func GetPokedexByGame(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	gameID := c.Params("game_id")

	var trainer models.Trainer
	if err := database.DB.Where("user_id = ?", userID).First(&trainer).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Trainer not found")
	}

	var entries []models.TrainerPokedexEntry
	if err := database.DB.Where("trainer_id = ? AND game_id = ?", trainer.ID, gameID).Find(&entries).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not fetch pokedex entries")
	}

	return c.JSON(entries)
}

// POST /me/pokedex/:game_id/:pokemon_id
func UpsertPokedexEntry(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	gameID := c.Params("game_id")
	pokemonID := c.Params("pokemon_id")

	var input PokedexEntryInput
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	var trainer models.Trainer
	if err := database.DB.Where("user_id = ?", userID).First(&trainer).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Trainer not found")
	}

	var entry models.TrainerPokedexEntry
	tx := database.DB.Where("trainer_id = ? AND game_id = ? AND pokemon_id = ?", trainer.ID, gameID, pokemonID).First(&entry)

	if tx.RowsAffected == 0 {
		entry = models.TrainerPokedexEntry{
			TrainerID:       trainer.ID,
			GameID:          uuid.MustParse(gameID),
			PokemonID:       uuid.MustParse(pokemonID),
			Caught:          input.Caught,
			Shiny:           input.Shiny,
			LivingDex:       input.LivingDex,
			ShinyLivingDex:  input.ShinyLivingDex,
			Notes:           input.Notes,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}
		if err := database.DB.Create(&entry).Error; err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Could not create entry")
		}
	} else {
		entry.Caught = input.Caught
		entry.Shiny = input.Shiny
		entry.LivingDex = input.LivingDex
		entry.ShinyLivingDex = input.ShinyLivingDex
		entry.Notes = input.Notes
		entry.UpdatedAt = time.Now()
		if err := database.DB.Save(&entry).Error; err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Could not update entry")
		}
	}

	return c.JSON(entry)
}

// PUT /me/pokedex/entry/:entry_id
func UpdatePokedexEntry(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	entryID := c.Params("entry_id")

	var input PokedexEntryInput
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	var trainer models.Trainer
	if err := database.DB.Where("user_id = ?", userID).First(&trainer).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Trainer not found")
	}

	var entry models.TrainerPokedexEntry
	if err := database.DB.Where("id = ? AND trainer_id = ?", entryID, trainer.ID).First(&entry).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Entry not found")
	}

	entry.Caught = input.Caught
	entry.Shiny = input.Shiny
	entry.LivingDex = input.LivingDex
	entry.ShinyLivingDex = input.ShinyLivingDex
	entry.Notes = input.Notes
	entry.UpdatedAt = time.Now()

	if err := database.DB.Save(&entry).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update entry")
	}

	return c.JSON(entry)
}

// DELETE /me/pokedex/entry/:entry_id
func DeletePokedexEntry(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	entryID := c.Params("entry_id")

	var trainer models.Trainer
	if err := database.DB.Where("user_id = ?", userID).First(&trainer).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Trainer not found")
	}

	if err := database.DB.Where("id = ? AND trainer_id = ?", entryID, trainer.ID).Delete(&models.TrainerPokedexEntry{}).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete entry")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
