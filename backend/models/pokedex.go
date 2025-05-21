package models

import (
	"time"

	"github.com/google/uuid"
)

type TrainerPokedexEntry struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	TrainerID  uuid.UUID `gorm:"type:uuid;not null"`
	GameID     uuid.UUID `gorm:"type:uuid;not null"`
	PokemonID  uuid.UUID `gorm:"type:uuid;not null"`

	Caught          bool `gorm:"default:false"`
	Shiny           bool `gorm:"default:false"`
	LivingDex       bool `gorm:"default:false"`
	ShinyLivingDex  bool `gorm:"default:false"`

	Notes      string    `gorm:"type:varchar(50)"`

	UpdatedAt  time.Time
	CreatedAt  time.Time
}
