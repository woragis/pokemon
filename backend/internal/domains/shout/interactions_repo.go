package shout

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type iRepository interface {
	// Count
	getInteractions(shoutID uuid.UUID) (*shoutInteractions, error)

	// Likes
	createLike(shoutLike *ShoutLike) error
	deleteLike(shoutID, userID uuid.UUID) error

	// Comments
	createComment(comment *ShoutComment) error
	updateComment(comment *ShoutComment) error
	deleteComment(commentID uuid.UUID) error

	// Views
	createView(view *ShoutView) error

	// Views
	createSave(save *ShoutSave) error
	deleteSave(shoutID, userID uuid.UUID) error
}

type iRepo struct {
	db *gorm.DB
}

func newInteractionRepo(db *gorm.DB) iRepository {
	return &iRepo{db: db}
}

/********************
 * GET INTERACTIONS *
 ********************/
type shoutInteractions struct {
	LikeCount    int64
	CommentCount int64
	RetweetCount int64
	ViewCount    int64
	SaveCount    int64
}

func (r *iRepo) getInteractions(shoutID uuid.UUID) (*shoutInteractions, error) {
	var result shoutInteractions

	// Count likes
	if err := r.db.Model(&ShoutLike{}).Where("shout_id = ?", shoutID).Count(&result.LikeCount).Error; err != nil {
		return nil, err
	}

	// Count comments
	if err := r.db.Model(&ShoutComment{}).Where("shout_id = ?", shoutID).Count(&result.CommentCount).Error; err != nil {
		return nil, err
	}

	// Count retweets (reshouts)
	if err := r.db.Model(&Shout{}).Where("reshout_of_id = ?", shoutID).Count(&result.RetweetCount).Error; err != nil {
		return nil, err
	}

	// Count views
	if err := r.db.Model(&ShoutView{}).Where("shout_id = ?", shoutID).Count(&result.ViewCount).Error; err != nil {
		return nil, err
	}

	// Count saves
	if err := r.db.Model(&ShoutSave{}).Where("shout_id = ?", shoutID).Count(&result.SaveCount).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

/********
 * LIKE *
 ********/

func (r* iRepo) createLike(shoutLike *ShoutLike) error {
	return r.db.Create(shoutLike).Error
}

func (r *iRepo) deleteLike(shoutID, userID uuid.UUID) error {
	return r.db.Delete(&ShoutLike{}, "shout_id = ? AND user_id = ?", shoutID, userID).Error
}

/************
 * Comments *
 ************/

func (r *iRepo) createComment(comment *ShoutComment) error {
	return r.db.Create(comment).Error
}

func (r *iRepo) updateComment(comment *ShoutComment) error {
	return r.db.Save(comment).Error
}

func (r *iRepo) deleteComment(commentID uuid.UUID) error {
	return r.db.Delete(&ShoutComment{}, "id = ?", commentID).Error
}

/*********
 * VIEWS *
 *********/

func (r *iRepo) createView(view *ShoutView) error {
	return r.db.Create(view).Error
}

/********
 * SAVE *
 *********/

func (r *iRepo) createSave(save *ShoutSave) error {
	return r.db.Create(save).Error
}

func (r *iRepo) deleteSave(shoutID, userID uuid.UUID) error {
	return r.db.Delete(&ShoutSave{}, "shout_id = ? AND user_id", shoutID, userID).Error
}
