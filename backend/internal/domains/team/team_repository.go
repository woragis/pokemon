package team

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/************************
 * REPOSITORY INTERFACE *
 ************************/

type TeamRepository interface {
	Create(team *Team) error
	GetByID(id uuid.UUID) (*Team, error)
	ListByUser(userID uuid.UUID, limit int, offset int) ([]Team, error)
	Update(team *Team) error
	Delete(id uuid.UUID) error
}

/*****************************
 * REPOSITORY IMPLEMENTATION *
 *****************************/

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepository{db}
}

func (r *teamRepository) Create(team *Team) error {
	return r.db.Create(team).Error
}

func (r *teamRepository) GetByID(id uuid.UUID) (*Team, error) {
	var team Team
	err := r.db.Preload("Pokemon").First(&team, "id = ?", id).Error
	return &team, err
}

func (r *teamRepository) ListByUser(userID uuid.UUID, limit int, offset int) ([]Team, error) {
	var teams []Team
	err := r.db.
		Preload("Pokemon").
		Limit(limit).
		Offset(offset).
		Where("user_id = ?", userID).
		Find(&teams).Error
	return teams, err
}

func (r *teamRepository) Update(team *Team) error {
	return r.db.Save(team).Error
}

func (r *teamRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Team{}, "id = ?", id).Error
}
