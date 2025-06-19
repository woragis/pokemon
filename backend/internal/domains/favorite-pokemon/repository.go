package favoritepokemon

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/************************
 * REPOSITORY INTERFACE *
 ************************/

type favoritepokemonRepository interface {
	create(mon *FavoritePokemon) error
	listByPopular() ([]FavoritePokemon, error)
	listByUser(userID uuid.UUID) ([]FavoritePokemon, error)
	getByID(id uuid.UUID) (*FavoritePokemon, error)
	update(mon *FavoritePokemon) error
	delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func newFavoritePokemonRepository(db *gorm.DB) favoritepokemonRepository {
	return &repository{db}
}

/*****************************
 * REPOSITORY IMPLEMENTATION *
 *****************************/

func (r *repository) create(mon *FavoritePokemon) error {
	return r.db.Create(mon).Error
}

func (r *repository) listByPopular() ([]FavoritePokemon, error) {
	var favorites []FavoritePokemon
	err := r.db.
		Order("like_count DESC").
		Find(&favorites).Error
	return favorites, err
}

func (r *repository) listByUser(userID uuid.UUID) ([]FavoritePokemon, error) {
	var favorites []FavoritePokemon
	err := r.db.
		Where("user_id = ?", userID).
		Find(&favorites).Error
	return favorites, err
}

func (r *repository) getByID(id uuid.UUID) (*FavoritePokemon, error) {
	var mon FavoritePokemon
	err := r.db.First(&mon, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &mon, nil
}

func (r *repository) update(mon *FavoritePokemon) error {
	return r.db.Save(mon).Error
}

func (r *repository) delete(id uuid.UUID) error {
	return r.db.Delete(&FavoritePokemon{}, "id = ?", id).Error
}
