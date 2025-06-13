package favoritepokemon

import "gorm.io/gorm"

type FavoritePokemonMigrator struct{}

func (m FavoritePokemonMigrator) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&FavoritePokemon{},
	)
}


