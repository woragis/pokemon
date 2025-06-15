package shout

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type shoutRepository interface {
	create(shout *Shout) error
	getById(id uuid.UUID) (*Shout, error)
	update(shout *Shout) error
	delete(id uuid.UUID) error

	list(limit, offset int) ([]Shout, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]Shout, error)
	listByParent(reshoutID uuid.UUID, limit, offset int) ([]Shout, error)
	countByUser(userID uuid.UUID) (int64, error)
}

type shoutRepo struct {
	db *gorm.DB
}

func newShoutRepo(db *gorm.DB) shoutRepository {
	return &shoutRepo{db}
}

func (r *shoutRepo) create(shout *Shout) error {
	return r.db.Create(shout).Error
}

func (r *shoutRepo) getById(id uuid.UUID) (*Shout, error) {
	var shout Shout
	err := r.db.
		Preload("User").
		Preload("Likes").
		Preload("Comments").
		Preload("ReshoutOf").
		Preload("QuoteComment").
		First(&shout, "id = ?", id).Error
	return &shout, err
}

func (r *shoutRepo) update(shout *Shout) error {
	return r.db.Save(shout).Error
}

func (r *shoutRepo) delete(id uuid.UUID) error {
	return r.db.Delete(&Shout{}, "id = ?", id).Error
}

func (r *shoutRepo) list(limit, offset int) ([]Shout, error) {
	var shouts []Shout
	err := r.db.
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&shouts).Error
	return shouts, err
}

func (r *shoutRepo) listByUser(userID uuid.UUID, limit, offset int) ([]Shout, error) {
	var shouts []Shout
	err := r.db.
		Preload("User").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&shouts).Error
	return shouts, err
}

func (r *shoutRepo) listByParent(reshoutID uuid.UUID, limit, offset int) ([]Shout, error) {
	var shouts []Shout
	err := r.db.
		Preload("User").
		Where("reshout_of_id = ?", reshoutID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&shouts).Error
	return shouts, err
}

func (r *shoutRepo) countByUser(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&Shout{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}
