package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatValues struct {
	HP   int `json:"hp"`
	Atk  int `json:"atk"`
	Def  int `json:"def"`
	SpA  int `json:"spa"`
	SpD  int `json:"spd"`
	Spe  int `json:"spe"`
}

type Move struct {
	ID int `json:"id"`
}

type Moves []Move

type Team struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Public      bool           `gorm:"default:true" json:"public"`
	Pokemon     []PokemonSlot  `gorm:"foreignKey:TeamID" json:"pokemon"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type PokemonSlot struct {
	ID          uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TeamID      uuid.UUID   `gorm:"type:uuid;not null;index" json:"team_id"`

	Slot        int         `gorm:"not null" json:"slot"`         // 1-6
	PokemonID   int         `gorm:"not null" json:"pokemon_id"`   // from PokeAPI species
	PokemonName string      `gorm:"not null" json:"pokemon_name"` // display convenience

	Level       int         `gorm:"default:50" json:"level"`      // default level 50
	NatureID    int         `json:"nature_id"`                    // PokeAPI ID
	GenderID    int         `json:"gender_id"`                    // PokeAPI ID (1 = female, 2 = male, 3 = genderless)
	AbilityID   int         `json:"ability_id"`                   // PokeAPI ID
	ItemID      int         `json:"item_id"`                      // PokeAPI ID

	Moves       Moves       `gorm:"type:jsonb" json:"moves"`      // slice of move IDs
	IVs         StatValues  `gorm:"type:jsonb" json:"ivs"`        // individual values
	EVs         StatValues  `gorm:"type:jsonb" json:"evs"`        // effort values
}

/****************
 * INTERACTIONS *
 ****************/

// --- Team Like ---
type TeamLike struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_team_like" json:"user_id"`
	TeamID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_team_like" json:"team_id"`
}

// --- Team View ---
type TeamView struct {
	ID       uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID   *uuid.UUID `gorm:"type:uuid;index" json:"user_id,omitempty"` // optional
	TeamID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"team_id"`
	ViewedAt time.Time  `gorm:"autoCreateTime" json:"viewed_at"`
}

// --- Team Save ---
type TeamSave struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_team_save" json:"user_id"`
	TeamID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_user_team_save" json:"team_id"`
}

// --- Team Comment with nesting support ---
type TeamComment struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	TeamID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"team_id"`
	ParentID  *uuid.UUID     `gorm:"type:uuid;index" json:"parent_id,omitempty"` // nullable, for nested comments
	Content   string         `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
