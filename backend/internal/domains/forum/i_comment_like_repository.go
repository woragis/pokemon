package forum

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*************************************
 * COMMENT LIKE REPOSITORY INTERFACE *
 *************************************/

type commentLikeRepository interface {
	create(like *CommentLike) error
	update(like *CommentLike) error
	get(commentID, userID uuid.UUID) (*CommentLike, error)
	delete(commentID, userID uuid.UUID) error
	count(commentID uuid.UUID) (likes, dislikes int64, err error)
}

type commentLikeRepo struct {
	db *gorm.DB
}

// Constructor
func newCommentLikeRepository(db *gorm.DB) commentLikeRepository {
	return &commentLikeRepo{db: db}
}

/******************************************
 * COMMENT LIKE REPOSITORY IMPLEMENTATION *
 ******************************************/

// Create a new CommentLike
func (r *commentLikeRepo) create(like *CommentLike) error {
	return r.db.Create(like).Error
}

// Update an existing CommentLike
func (r *commentLikeRepo) update(like *CommentLike) error {
	return r.db.Save(like).Error
}

// Get retrieves a CommentLike by commentID and userID
func (r *commentLikeRepo) get(commentID, userID uuid.UUID) (*CommentLike, error) {
	var like CommentLike
	err := r.db.
		Where("comment_id = ? AND user_id = ?", commentID, userID).
		First(&like).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &like, err
}

// Delete removes a CommentLike by commentID and userID
func (r *commentLikeRepo) delete(commentID, userID uuid.UUID) error {
	return r.db.
		Where("comment_id = ? AND user_id = ?", commentID, userID).
		Delete(&CommentLike{}).Error
}

func (r *commentLikeRepo) count(commentID uuid.UUID) (int64, int64, error) {
	var likes int64
	var dislikes int64

	err := r.db.Model(&CommentLike{}).
		Where("comment_id = ? AND like = TRUE", commentID).
		Count(&likes).Error
	if err != nil {
		return 0, 0, err
	}

	err = r.db.Model(&CommentLike{}).
		Where("comment_id = ? AND like = FALSE", commentID).
		Count(&dislikes).Error
	if err != nil {
		return 0, 0, err
	}

	return likes, dislikes, nil
}
