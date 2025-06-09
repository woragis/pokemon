package user

import (
	"gorm.io/gorm"
)

type UserMigrator struct{}

func (m UserMigrator) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
	)
}
