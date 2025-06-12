package forum

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*********************************
 * CATEGORY REPOSITORY INTERFACE *
 *********************************/

type topicCategoryRepository interface {
	create(c *TopicCategory) error
	getByID(id uuid.UUID) (*TopicCategory, error)
	update(c *TopicCategory) error
	delete(id uuid.UUID) error
	list(limit, offset int) ([]TopicCategory, error)
}

type topicCategoryRepositoryImpl struct {
	db *gorm.DB
}

func newTopicCategoryRepository(db *gorm.DB) topicCategoryRepository {
	return &topicCategoryRepositoryImpl{db: db}
}

/**************************************
 * CATEGORY REPOSITORY IMPLEMENTATION *
 **************************************/

func (r *topicCategoryRepositoryImpl) create(c *TopicCategory) error {
	return r.db.Create(c).Error
}

func (r *topicCategoryRepositoryImpl) getByID(id uuid.UUID) (*TopicCategory, error) {
	var category TopicCategory
	err := r.db.First(&category, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *topicCategoryRepositoryImpl) update(c *TopicCategory) error {
	return r.db.Save(c).Error
}

func (r *topicCategoryRepositoryImpl) delete(id uuid.UUID) error {
	return r.db.Delete(&TopicCategory{}, "id = ?", id).Error
}

func (r *topicCategoryRepositoryImpl) list(limit, offset int) ([]TopicCategory, error) {
	var categories []TopicCategory
	err := r.db.
		Order("name ASC").
		Limit(limit).
		Offset(offset).
		Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
