package blog

import "gorm.io/gorm"

type BlogMigrator struct{}

func (m BlogMigrator) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Post{},
		&PostLike{},
		&PostComment{},
		&PostView{},
		&PostSave{},
	)
}
