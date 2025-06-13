package forum

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*****************************
 * LIKE REPOSITORY INTERFACE *
 *****************************/

type topicLikeRepository interface {
	create(like *TopicLike) error
	get(topicID, userID uuid.UUID) (*TopicLike, error)
	delete(topicID, userID uuid.UUID) error
	count(topicID uuid.UUID, likeValue bool) (int64, error)
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

func (r *topicLikeRepo) count(topicID uuid.UUID, likeValue bool) (int64, error) {
	var count int64
	err := r.db.Model(&TopicLike{}).
		Where("topic_id = ? AND like = ?", topicID, likeValue).
		Count(&count).Error
	return count, err
}

func (r *topicLikeRepo) create(like *TopicLike) error {
	// Delete any existing like for the user-topic pair (if exists)
	if err := r.delete(like.TopicID, like.UserID); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Create the new like
	return r.db.Create(like).Error
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