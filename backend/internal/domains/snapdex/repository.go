package snapdex

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/************************
 * REPOSITORY INTERFACE *
 ************************/

type snapRepository interface {
	create(snap *Snap) error
	getById(id uuid.UUID) (*Snap, error)
	update(snap *Snap) error
	delete(id uuid.UUID) error

	list(limit, offset int) ([]Snap, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]Snap, error)
	countByUser(userID uuid.UUID) (int64, error)
}

type snapRepo struct {
	db *gorm.DB
}

func newRepo(db *gorm.DB) snapRepository {
	return &snapRepo{db: db}
}

/*****************************
 * REPOSITORY IMPLEMENTATION *
 *****************************/

func (r *snapRepo) create(snap *Snap) error {
	return r.db.Create(snap).Error
}

func (r *snapRepo) update(snap *Snap) error {
	return r.db.Save(snap).Error
}

func (r *snapRepo) delete(id uuid.UUID) error {
	return r.db.Delete(&Snap{}, "id = ?", id).Error
}

func (r *snapRepo) getById(id uuid.UUID) (*Snap, error) {
	var snap Snap
	err := r.db.Where("id = ?", id).First(&snap).Error // 🔧 fix: chain Where before First
	return &snap, err
}

func (r *snapRepo) list(limit, offset int) ([]Snap, error) {
	var snaps []Snap
	err := r.db.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&snaps).
		Error // 🔧 fix: Order, Limit, Offset must come before Find
	return snaps, err
}

func (r *snapRepo) listByUser(userID uuid.UUID, limit, offset int) ([]Snap, error) {
	var snaps []Snap
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&snaps).
		Error
	return snaps, err
}

func (r *snapRepo) countByUser(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.
		Model(&Snap{}). // 🔧 fix: use Model, not Find
		Where("user_id = ?", userID).
		Count(&count).
		Error
	return count, err
}
