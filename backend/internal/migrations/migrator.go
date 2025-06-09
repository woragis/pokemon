package migrations

import "gorm.io/gorm"

type Migrator interface {
	Migrate(db *gorm.DB) error
}
