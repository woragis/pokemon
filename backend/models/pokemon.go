package models

import "github.com/google/uuid"

type PokemonSpecies struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;not null"`  // e.g., "Bulbasaur"
	NationalDex int       `gorm:"not null"`         // e.g., 1
	TypePrimary string    `gorm:"not null"`         // e.g., "Grass"
	TypeSecondary string  `gorm:"default:null"`     // e.g., "Poison"

	// Optional fields
	BaseEvolution string  `gorm:"default:null"`     // "Bulbasaur"
	FinalEvolution string `gorm:"default:null"`     // "Venusaur"
}
