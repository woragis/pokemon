package forum

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*****************************
 * VIEW REPOSITORY INTERFACE *
 *****************************/

type topicViewRepository interface {
	create(view *TopicView) error
	listByUser(userID uuid.UUID, limit, offset int) ([]TopicView, error)
}

type topicViewRepoImpl struct {
	db *gorm.DB
}

func newTopicViewRepository(db *gorm.DB) topicViewRepository {
	return &topicViewRepoImpl{db: db}
}

/**********************************
 * VIEW REPOSITORY IMPLEMENTATION *
 **********************************/

func (r *topicViewRepoImpl) create(view *TopicView) error {
	return r.db.Create(view).Error
}

func (r *topicViewRepoImpl) listByUser(userID uuid.UUID, limit, offset int) ([]TopicView, error) {
	var views []TopicView
	err := r.db.
		Preload("Topic").
		Preload("User").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&views).Error

	return views, err
}