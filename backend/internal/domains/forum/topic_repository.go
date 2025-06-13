package forum

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/************************
 * REPOSITORY INTERFACE *
 ************************/

type topicRepository interface {
	create(topic *Topic) error
	getByID(id uuid.UUID) (*Topic, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]Topic, error)
	list(limit, offset int) ([]Topic, error)
	update(topic *Topic) error
	delete(id uuid.UUID) error
}

/*****************************
 * REPOSITORY IMPLEMENTATION *
 *****************************/

type repository struct {
	db *gorm.DB
}

func newTopicRepository(db *gorm.DB) topicRepository {
	return &repository{db}
}

func (r *repository) create(topic *Topic) error {
	return r.db.Create(topic).Error
}

func (r *repository) getByID(id uuid.UUID) (*Topic, error) {
	var topic Topic
	err := r.db.Preload("User").First(&topic, "id = ?", id).Error
	return &topic, err
}

func (r *repository) listByUser(userID uuid.UUID, limit, offset int) ([]Topic, error) {
	var topics []Topic
	err := r.db.
		Preload("User").
		Order("created_at DESC").
		Where("user_id= ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&topics).Error
	return topics, err
}

func (r *repository) list(limit, offset int) ([]Topic, error) {
	var topics []Topic
	err := r.db.
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&topics).Error
	return topics, err
}

func (r *repository) update(topic *Topic) error {
	return r.db.Save(topic).Error
}

func (r *repository) delete(id uuid.UUID) error {
	return r.db.Delete(&Topic{}, "id = ?", id).Error
}
