package blog

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/******************************
 ******************************
 ************ MAIN ************
 ******************************
 ******************************/

/************************
 * REPOSITORY INTERFACE *
 ************************/

type blogRepository interface {
	list(limit int, offset int) ([]Post, error)
	listByUser(userID uuid.UUID, limit int, offset int) ([]Post, error)
	// countByUser(userID uuid.UUID) (int64, error)

	create(post *Post) error
	getByID(id uuid.UUID) (*Post, error)
	update(post *Post) error
	delete(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func newRepository(db *gorm.DB) blogRepository {
	return &repository{db}
}

/*****************************
 * REPOSITORY IMPLEMENTATION *
 *****************************/

func (r *repository) create(post *Post) error {
	return r.db.Create(post).Error
}

func (r *repository) getByID(id uuid.UUID) (*Post, error) {
	var post Post
	err := r.db.Preload("Userr").First(&post, "id = ?", id).Error
	return &post, err
}

func (r *repository) listByUser(userID uuid.UUID, limit int, offset int) ([]Post, error) {
	var posts []Post
	err := r.db.
		Preload("Userr").
		Limit(limit).
		Offset(offset).
		Where("user_id = ?", userID).
		Find(&posts).Error
	return posts, err
}

func (r *repository) list(limit int, offset int) ([]Post, error) {
	var posts []Post
	err := r.db.
		Preload("User").
		Limit(limit).
		Offset(offset).
		Order("created_at ASC").
		Find(&posts).Error
	return posts, err
}

func (r *repository) update(post *Post) error {
	return r.db.Save(post).Error
}

func (r *repository) delete(id uuid.UUID) error {
	return r.db.Delete(&Post{}, "id = ?", id).Error
}

/**************************************
 **************************************
 ************ INTERACTIONS ************
 **************************************
 **************************************/
