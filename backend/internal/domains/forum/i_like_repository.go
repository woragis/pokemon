package forum

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*****************************
 * LIKE REPOSITORY INTERFACE *
 *****************************/

type topicLikeRepository interface {
	create(like *TopicLike) error
	update(like *TopicLike) error
	get(topicID, userID uuid.UUID) (*TopicLike, error)
	delete(topicID, userID uuid.UUID) error
}

type topicLikeRepo struct {
	db *gorm.DB
}

func newTopicLikeRepository(db *gorm.DB) topicLikeRepository {
	return &topicLikeRepo{db: db}
}

/**********************************
 * LIKE REPOSITORY IMPLEMENTATION *
 **********************************/

func (r *topicLikeRepo) create(like *TopicLike) error {
	return r.db.Create(like).Error
}

func (r *topicLikeRepo) update(like *TopicLike) error {
	return r.db.Save(like).Error
}

func (r *topicLikeRepo) get(topicID, userID uuid.UUID) (*TopicLike, error) {
	var like TopicLike
	err := r.db.Where("topic_id = ? AND user_id = ?", topicID, userID).First(&like).Error
	if err != nil {
		return nil, err
	}
	return &like, nil
}

func (r *topicLikeRepo) delete(topicID, userID uuid.UUID) error {
	return r.db.Where("topic_id = ? AND user_id = ?", topicID, userID).Delete(&TopicLike{}).Error
}