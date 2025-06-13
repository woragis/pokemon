package team

import "gorm.io/gorm"

type TeamMigrator struct{}

func (m TeamMigrator) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Team{},
		&PokemonSlot{},
		&TeamLike{},
		&TeamView{},
		&TeamSave{},
		&TeamComment{},
	)
}


