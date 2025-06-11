package team

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/************************
 * REPOSITORY INTERFACE *
 ************************/

type teamRepository interface {
	create(team *Team) error
	getByID(id uuid.UUID) (*Team, error)
	listByUser(userID uuid.UUID, limit int, offset int) ([]Team, error)
	update(team *Team) error
	delete(id uuid.UUID) error
}

/*****************************
 * REPOSITORY IMPLEMENTATION *
 *****************************/

type repository struct {
	db *gorm.DB
}

func newRepository(db *gorm.DB) teamRepository {
	return &repository{db}
}

func (r *repository) create(team *Team) error {
	return r.db.Create(team).Error
}

func (r *repository) getByID(id uuid.UUID) (*Team, error) {
	var team Team
	err := r.db.Preload("Pokemon").First(&team, "id = ?", id).Error
	return &team, err
}

func (r *repository) listByUser(userID uuid.UUID, limit int, offset int) ([]Team, error) {
	var teams []Team
	err := r.db.
		Preload("Pokemon").
		Limit(limit).
		Offset(offset).
		Where("user_id = ?", userID).
		Find(&teams).Error
	return teams, err
}

func (r *repository) update(team *Team) error {
	return r.db.Save(team).Error
}

func (r *repository) delete(id uuid.UUID) error {
	return r.db.Delete(&Team{}, "id = ?", id).Error
}
