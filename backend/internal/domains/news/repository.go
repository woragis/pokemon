package news

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/**********************
 **********************
 ******** MAIN ********
 **********************
 **********************/

type repository interface {
	list(limit, offset int) ([]News, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]News, error)
	countByUser(userID uuid.UUID) (int64, error)

	get(ID uuid.UUID) (News, error)
	create(news News) error
	update(news News) error
	delete(ID, userID uuid.UUID) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func newRepo(db *gorm.DB) repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) list(limit, offset int) ([]News, error) {
	var news []News
	err := r.db.
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&news).Error
	return news, err
}

func (r *repositoryImpl) listByUser(userID uuid.UUID, limit, offset int) ([]News, error) {
	var news []News
	err := r.db.
		Where("user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&news).Error
	return news, err
}

func (r *repositoryImpl) get(ID uuid.UUID) (News, error) {
	var n News
	err := r.db.
		First(&n, "id = ?", ID).Error
	return n, err
}

func (r *repositoryImpl) create(news News) error {
	return r.db.Create(&news).Error
}

func (r *repositoryImpl) update(news News) error {
	return r.db.Save(&news).Error
}

func (r *repositoryImpl) delete(ID, userID uuid.UUID) error {
	return r.db.Delete(&News{}, "id = ? AND user_id = ?", ID, userID).Error
}

func (r *repositoryImpl) countByUser(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&News{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

/******************************
 ******************************
 ******** INTERACTIONS ********
 ******************************
 ******************************/
