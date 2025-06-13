package forum

import "gorm.io/gorm"

type ForumMigrator struct{}

func (m ForumMigrator) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Topic{},
		&TopicView{},
		&TopicLike{},
		&TopicComment{},
		&TopicCommentLike{},
		&TopicCategory{},
	)
}

