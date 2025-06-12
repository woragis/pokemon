package forum

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/********************************
 * COMMENT REPOSITORY INTERFACE *
 ********************************/

type topicCommentRepository interface {
	create(c *TopicComment) error
	getByID(id uuid.UUID) (*TopicComment, error)
	update(c *TopicComment) error
	delete(id uuid.UUID) error
	listByTopic(topicID uuid.UUID) ([]TopicComment, error)
	count(topicID uuid.UUID) (int64, error)
}

type topicCommentRepositoryImpl struct {
	db *gorm.DB
}

func newTopicCommentRepository(db *gorm.DB) topicCommentRepository {
	return &topicCommentRepositoryImpl{db: db}
}

/*************************************
 * COMMENT REPOSITORY IMPLEMENTATION *
 *************************************/

func (r *topicCommentRepositoryImpl) create(c *TopicComment) error {
	return r.db.Create(c).Error
}

func (r *topicCommentRepositoryImpl) getByID(id uuid.UUID) (*TopicComment, error) {
	var comment TopicComment
	if err := r.db.Preload("User").Preload("Replies").First(&comment, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *topicCommentRepositoryImpl) update(c *TopicComment) error {
	return r.db.Save(c).Error
}

func (r *topicCommentRepositoryImpl) delete(id uuid.UUID) error {
	return r.db.Delete(&TopicComment{}, "id = ?", id).Error
}

func (r *topicCommentRepositoryImpl) listByTopic(topicID uuid.UUID) ([]TopicComment, error) {
	var comments []TopicComment
	err := r.db.
		Preload("User").
		Preload("Replies").
		Where("topic_id = ? AND parent_id IS NULL", topicID).
		Order("created_at ASC").
		Find(&comments).Error
	return comments, err
}

func (r *topicCommentRepositoryImpl) count(topicID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.
		Where("topic_id = ?", topicID).
		Count(&count).Error
	return count, err
}
