package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	Name      string         `gorm:"not null" json:"name"`
	Description string       `json:"description"`
	Public    bool           `gorm:"default:true" json:"public"`
	Pokemon   []PokemonSlot  `gorm:"foreignKey:TeamID" json:"pokemon"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type PokemonSlot struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TeamID   uuid.UUID `gorm:"type:uuid;not null;index" json:"team_id"`

	Slot     int        `gorm:"not null" json:"slot"` // 1-6
	PokemonName string  `gorm:"not null" json:"pokemon_name"`
	Level     int       `json:"level"`
	Ability   string    `json:"ability"`
	Item      string    `json:"item"`
	Moves     []string  `gorm:"type:text[]" json:"moves"` // Or serialize as JSON if needed

	IVs       string    `json:"ivs"`  // can also be a struct if detailed
	EVs       string    `json:"evs"`

	// Optional extra fields like Nature, Gender, etc.
}

/****************
 * INTERACTIONS *
 ****************/

type TeamLike struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	TeamID uuid.UUID `gorm:"type:uuid;not null;index" json:"team_id"`
}

type TeamView struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;index" json:"user_id"` // optional if anonymous
	TeamID    uuid.UUID `gorm:"type:uuid;not null;index" json:"team_id"`
	ViewedAt  time.Time `gorm:"autoCreateTime" json:"viewed_at"`
}

type TeamSave struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	TeamID uuid.UUID `gorm:"type:uuid;not null;index" json:"team_id"`
}

type TeamComment struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	TeamID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"team_id"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
