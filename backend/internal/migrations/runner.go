package migrations

import (
	"gorm.io/gorm"
)

func RunAll(db *gorm.DB, migrators []Migrator) error {
	for _, m := range migrators {
		if err := m.Migrate(db); err != nil {
			return err
		}
	}
	return nil
}
