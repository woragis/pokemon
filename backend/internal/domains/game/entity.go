package game

import "github.com/google/uuid"

type PokemonGame struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"unique;not null" json:"name"`
	Generation  int       `gorm:"not null" json:"generation"`       // e.g., 1, 2, 3, etc.
	ReleasedAt  string    `gorm:"not null" json:"released_at"`       // ISO or "YYYY-MM-DD"
	Description string    `gorm:"type:text" json:"description"`
}
