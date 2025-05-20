package models

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;not null"`
	Region      string    `gorm:"not null"`       // e.g., "Kanto", "Hoenn"
	Generation  int       `gorm:"not null"`       // e.g., 1, 2, 3, etc.
	ReleasedAt  string    `gorm:"not null"`       // ISO or "YYYY-MM-DD"
	Description string    `gorm:"type:text"`
}

type UserPokedexEntry struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	TrainerID  uuid.UUID `gorm:"type:uuid;not null"`
	GameID     uuid.UUID `gorm:"type:uuid;not null"`
	PokemonID  uuid.UUID `gorm:"type:uuid;not null"`

	Seen       bool      `gorm:"default:false"`
	Caught     bool      `gorm:"default:false"`
	Evolved    bool      `gorm:"default:false"`
	Shiny      bool      `gorm:"default:false"`
	LivingDex  bool      `gorm:"default:false"`

	Form       *string   `gorm:"type:varchar(50)"`

	UpdatedAt  time.Time
	CreatedAt  time.Time

	Trainer    Trainer          `gorm:"foreignKey:TrainerID"`
	Game       Game             `gorm:"foreignKey:GameID"`
	Pokemon    PokemonSpecies   `gorm:"foreignKey:PokemonID"`
}
