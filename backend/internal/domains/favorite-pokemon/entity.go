package favoritepokemon

import (
	"errors"
	"pokemon/internal/domains/user"
	"strings"

	"github.com/google/uuid"
)

type FavoritePokemon struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User      user.User `json:"user" gorm:"foreignKey:UserID"`
	PokemonID int       `json:"pokemon_id" gorm:"type:int;not null"`
	Nickname  string    `json:"nickname" gorm:"type:varchar(20)"`
	LikeCount int       `json:"like_count" gorm:"type:int;default:0"`
}

func (f *FavoritePokemon) Validate() error {
	f.Nickname = strings.TrimSpace(f.Nickname)
	if f.PokemonID <= 0 {
		return errors.New("pokemon_id must be a positive integer")
	}
	if len(f.Nickname) > 20 {
		return errors.New("nickname cannot exceed 20 characters")
	}
	return nil
}
