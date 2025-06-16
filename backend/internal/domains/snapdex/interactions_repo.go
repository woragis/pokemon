package snapdex

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*************************************
 * INTERACTIONS REPOSITORY INTERFACE *
 *************************************/

type interactionRepository interface {
	// SnapComment
	createComment(comment *SnapComment) error
	deleteComment(id uuid.UUID) error
	listComments(snapID uuid.UUID, limit, offset int) ([]SnapComment, error)
	countComments(snapID uuid.UUID) (int64, error)

	// SnapLike
	createLike(like *SnapLike) error
	deleteLike(snapID, userID uuid.UUID) error
	hasLiked(snapID, userID uuid.UUID) (bool, error)
	countLikes(snapID uuid.UUID) (int64, error)

	// SnapCommentLike
	createCommentLike(like *SnapCommentLike) error
	deleteCommentLike(commentID, userID uuid.UUID) error
	hasCommentLiked(commentID, userID uuid.UUID) (bool, error)
	countCommentLikes(commentID uuid.UUID) (int64, error)
}

/******************************
 * REPOSITORY STRUCT & CONSTRUCTOR *
 ******************************/

type interactionRepo struct {
	db *gorm.DB
}

func newInteractionRepo(db *gorm.DB) interactionRepository {
	return &interactionRepo{db: db}
}

/******************************
 * SnapComment IMPLEMENTATION *
 ******************************/

func (r *interactionRepo) createComment(comment *SnapComment) error {
	comment.CreatedAt = time.Now()
	return r.db.Create(comment).Error
}

func (r *interactionRepo) deleteComment(id uuid.UUID) error {
	return r.db.Delete(&SnapComment{}, "id = ?", id).Error
}

func (r *interactionRepo) listComments(snapID uuid.UUID, limit, offset int) ([]SnapComment, error) {
	var comments []SnapComment
	err := r.db.
		Where("snap_id = ?", snapID).
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(&comments).
		Error
	return comments, err
}

func (r *interactionRepo) countComments(snapID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.
		Model(&SnapComment{}).
		Where("snap_id = ?", snapID).
		Count(&count).
		Error
	return count, err
}

/******************************
 * SnapLike IMPLEMENTATION *
 ******************************/

func (r *interactionRepo) createLike(like *SnapLike) error {
	like.CreatedAt = time.Now()
	return r.db.Create(like).Error
}

func (r *interactionRepo) deleteLike(snapID, userID uuid.UUID) error {
	return r.db.
		Where("snap_id = ? AND user_id = ?", snapID, userID).
		Delete(&SnapLike{}).
		Error
}

func (r *interactionRepo) hasLiked(snapID, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.
		Model(&SnapLike{}).
		Where("snap_id = ? AND user_id = ?", snapID, userID).
		Count(&count).
		Error
	return count > 0, err
}

func (r *interactionRepo) countLikes(snapID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.
		Model(&SnapLike{}).
		Where("snap_id = ?", snapID).
		Count(&count).
		Error
	return count, err
}

/******************************
 * SnapCommentLike IMPLEMENTATION *
 ******************************/

func (r *interactionRepo) createCommentLike(like *SnapCommentLike) error {
	like.CreatedAt = time.Now()
	return r.db.Create(like).Error
}

func (r *interactionRepo) deleteCommentLike(commentID, userID uuid.UUID) error {
	return r.db.
		Where("comment_id = ? AND user_id = ?", commentID, userID).
		Delete(&SnapCommentLike{}).
		Error
}

func (r *interactionRepo) hasCommentLiked(commentID, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.
		Model(&SnapCommentLike{}).
		Where("comment_id = ? AND user_id = ?", commentID, userID).
		Count(&count).
		Error
	return count > 0, err
}

func (r *interactionRepo) countCommentLikes(commentID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.
		Model(&SnapCommentLike{}).
		Where("comment_id = ?", commentID).
		Count(&count).
		Error
	return count, err
}
