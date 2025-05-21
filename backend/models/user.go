package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Username string            `gorm:"unique;not null" json:"username"`
	Name     string            `json:"name"`
	Email    string            `gorm:"unique;not null" json:"email"`
	Password string            `gorm:"not null" json:"-"`
	Role     string            `gorm:"default:user" json:"role"`

	FavoritePokemon string     `json:"favorite_pokemon"`

	FavoriteGameID uuid.UUID   `gorm:"type:uuid" json:"favorite_game_id"`
	FavoriteGame   PokemonGame `gorm:"foreignKey:FavoriteGameID" json:"favorite_game"`

	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"deleted_at,omitempty"`
}

type Trainer struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID   uuid.UUID `gorm:"type:uuid;not null;unique"`
	Nickname string    `gorm:"not null"`
	Avatar   string

	User     User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
