package migrations

import (
	favoritepokemon "pokemon/internal/domains/favorite-pokemon"
	"pokemon/internal/domains/forum"
	"pokemon/internal/domains/team"
	"pokemon/internal/domains/user"
)

func GetAllMigrators() []Migrator {
	return []Migrator{
		user.UserMigrator{},
		forum.ForumMigrator{},
		team.TeamMigrator{},
		favoritepokemon.FavoritePokemonMigrator{},
	}
}
