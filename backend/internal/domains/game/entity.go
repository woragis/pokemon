package game

import (
	"pokemon/internal/domains/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Game struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name       string    `json:"name" gorm:"type:text;not null"`
	Description string    `gorm:"type:text" json:"description"`
	Generation  int       `gorm:"not null" json:"generation"`       // e.g., 1, 2, 3, etc.
	ReleasedAt  time.Time `gorm:"not null" json:"released_at"`       // ISO or "YYYY-MM-DD"

	// Inclusive start (e.g., 1)
	// Inclusive end (e.g., 151)
	DexStartID uint `json:"dex_start_id" gorm:"type:integer;not null"`
	DexEndID   uint `json:"dex_end_id" gorm:"type:integer;not null"`

	// Bitmask: true if Pokémon is present in game
	Pokedex    []byte `gorm:"type:bytea"`

	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type GamePokedex struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID         uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	GameID         uuid.UUID `json:"game_id" gorm:"type:uuid;not null"`
	User           user.User `json:"user" gorm:"foreignKey:UserID"`
	Game           Game      `json:"game" gorm:"foreignKey:GameID"`

	Seen           []byte `json:"seen" gorm:"type:bytea"`
	Captured       []byte `json:"captured" gorm:"type:bytea"`
	ShinySeen      []byte `json:"shiny_seen" gorm:"type:bytea"`
	ShinyCaptured  []byte `json:"shiny_captured" gorm:"type:bytea"`

	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

func dexBitmaskSize(startID, endID int) int {
	numBits := (endID - startID + 1)
	return (numBits + 7) / 8
}

func createGameBitmask(startID, endID int, availableIDs []int) []byte {
	size := dexBitmaskSize(startID, endID)
	data := make([]byte, size)

	for _, id := range availableIDs {
		if id >= startID && id <= endID {
			data = setBit(data, id-startID) // Offset-based bit position
		}
	}
	return data
}

func setDexBit(data []byte, dexStart, pokeID int) []byte {
	index := pokeID - dexStart
	return setBit(data, index)
}

func isDexBitSet(data []byte, dexStart, pokeID int) bool {
	index := pokeID - dexStart
	return isBitSet(data, index)
}

func makeRange(start, end int) []int {
	result := make([]int, end-start+1)
	for i := range result {
		result[i] = start + i
	}
	return result
}

func setBit(data []byte, index int) []byte {
	byteIndex := index / 8
	bitIndex := index % 8
	if byteIndex >= len(data) {
		newData := make([]byte, byteIndex+1)
		copy(newData, data)
		data = newData
	}
	data[byteIndex] |= 1 << bitIndex
	return data
}

func clearBit(data []byte, index int) []byte {
	byteIndex := index / 8
	bitIndex := index % 8
	if byteIndex < len(data) {
		data[byteIndex] &^= 1 << bitIndex
	}
	return data
}

func isBitSet(data []byte, index int) bool {
	byteIndex := index / 8
	bitIndex := index % 8
	if byteIndex >= len(data) {
		return false
	}
	return data[byteIndex]&(1<<bitIndex) != 0
}
