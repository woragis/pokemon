package team

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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

type Gender int

const (
	Female Gender = 1
	Male   Gender = 2
	Genderless Gender = 3
)

type Moves []Move


type Team struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`

	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description,omitempty" gorm:"type:text"`
	Public      bool           `json:"public" gorm:"default:true"`

	Pokemon     []PokemonSlot  `json:"pokemon" gorm:"foreignKey:TeamID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}


type PokemonSlot struct {
	ID          uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TeamID      uuid.UUID   `gorm:"type:uuid;not null;index" json:"team_id"`

	Slot        int         `gorm:"not null" json:"slot"`         // 1-6
	PokemonID   int         `gorm:"not null" json:"pokemon_id"`   // from PokeAPI species
	PokemonName string      `gorm:"not null" json:"pokemon_name"` // display convenience

	Level       int         `gorm:"default:50" json:"level"`      // default level 50
	NatureID    int         `json:"nature_id"`                    // PokeAPI ID
	GenderID    Gender      `json:"gender_id"`                    // PokeAPI ID (1 = female, 2 = male, 3 = genderless)
	AbilityID   int         `json:"ability_id"`                   // PokeAPI ID
	ItemID      int         `json:"item_id"`                      // PokeAPI ID

	MoveList    Moves       `gorm:"type:jsonb" json:"moves"`      // slice of move IDs
	IVs         StatValues  `gorm:"type:jsonb" json:"ivs"`        // individual values
	EVs         StatValues  `gorm:"type:jsonb" json:"evs"`        // effort values
}

/***************
 * VALIDATIONS *
 ***************/

func (m *StatValues) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), m)
}

func (m StatValues) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (s StatValues) Validate(isEV bool) error {
	max := 31
	totalLimit := -1

	if isEV {
		max = 252
		totalLimit = 510
	}

	stats := []int{s.HP, s.Atk, s.Def, s.SpA, s.SpD, s.Spe}

	for _, v := range stats {
		if v < 0 || v > max {
			return fmt.Errorf("stat value %d out of bounds (0–%d)", v, max)
		}
	}

	if isEV {
		sum := 0
		for _, v := range stats {
			sum += v
		}
		if sum > totalLimit {
			return fmt.Errorf("total EVs exceed limit (%d > 510)", sum)
		}
	}

	return nil
}

func (m *Moves) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), m)
}

func (m Moves) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m Moves) Validate() error {
	if len(m) > 4 {
		return fmt.Errorf("a Pokémon can only have up to 4 moves")
	}

	moveSet := make(map[int]struct{})
	for _, move := range m {
		if _, exists := moveSet[move.ID]; exists {
			return fmt.Errorf("duplicate move ID %d", move.ID)
		}
		moveSet[move.ID] = struct{}{}
	}

	return nil
}

func (t *Team) Validate() error {
	if len(t.Pokemon) == 0 || len(t.Pokemon) > 6 {
		return fmt.Errorf("team must have between 1 and 6 Pokémon")
	}

	slots := make(map[int]struct{})
	for i := range t.Pokemon {
		slot := t.Pokemon[i].Slot

		if _, exists := slots[slot]; exists {
			return fmt.Errorf("duplicate slot %d", slot)
		}
		slots[slot] = struct{}{}

		if err := t.Pokemon[i].Validate(); err != nil {
			return fmt.Errorf("slot %d: %w", slot, err)
		}
	}

	return nil
}

func (p *PokemonSlot) Validate() error {
	if p.Slot < 1 || p.Slot > 6 {
		return fmt.Errorf("slot must be between 1 and 6")
	}
	if p.Level < 1 || p.Level > 100 {
		return fmt.Errorf("level must be between 1 and 100")
	}

	if err := p.IVs.Validate(false); err != nil {
		return fmt.Errorf("invalid IVs: %w", err)
	}
	if err := p.EVs.Validate(true); err != nil {
		return fmt.Errorf("invalid EVs: %w", err)
	}
	if err := p.MoveList.Validate(); err != nil {
		return fmt.Errorf("invalid moves: %w", err)
	}

	if p.GenderID < 1 || p.GenderID > 3 {
		return fmt.Errorf("invalid gender ID %d", p.GenderID)
	}

	return nil
}
