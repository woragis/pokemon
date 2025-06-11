package team

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/************************
 * REPOSITORY INTERFACE *
 ************************/

type teamInteractionRepository interface {
	// Likes
	createLike(like *TeamLike) error
	deleteLike(userID, teamID uuid.UUID) error
	countLikes(teamID uuid.UUID) (int64, error)

	// Views
	createView(view *TeamView) error
	countViews(teamID uuid.UUID) (int64, error)

	// Saves
	createSave(save *TeamSave) error
	deleteSave(userID, teamID uuid.UUID) error

	// Comments
	createComment(comment *TeamComment) error
	updateComment(comment *TeamComment) error
	deleteComment(id uuid.UUID) error
	getComments(teamID uuid.UUID, limit int, offset int) ([]TeamComment, error)
	countComments(teamID uuid.UUID) (int64, error)
}

/*****************************
 * REPOSITORY IMPLEMENTATION *
 *****************************/

type interactionRepository struct {
	db *gorm.DB
}

func newInteractionRepository(db *gorm.DB) teamInteractionRepository {
	return &interactionRepository{db}
}

// --- Likes ---
func (r *interactionRepository) createLike(like *TeamLike) error {
	return r.db.Create(like).Error
}

func (r *interactionRepository) deleteLike(userID, teamID uuid.UUID) error {
	return r.db.Where("user_id = ? AND team_id = ?", userID, teamID).Delete(&TeamLike{}).Error
}

func (r *interactionRepository) countLikes(teamID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&TeamLike{}).Where("team_id = ?", teamID).Count(&count).Error
	return count, err
}

// --- Views ---
func (r *interactionRepository) createView(view *TeamView) error {
	return r.db.Create(view).Error
}

func (r *interactionRepository) countViews(teamID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&TeamView{}).Where("team_id = ?", teamID).Count(&count).Error
	return count, err
}

// --- Saves ---
func (r *interactionRepository) createSave(save *TeamSave) error {
	return r.db.Create(save).Error
}

func (r *interactionRepository) deleteSave(userID, teamID uuid.UUID) error {
	return r.db.Where("user_id = ? AND team_id = ?", userID, teamID).Delete(&TeamSave{}).Error
}

// --- Comments ---
func (r *interactionRepository) createComment(comment *TeamComment) error {
	return r.db.Create(comment).Error
}

func (r *interactionRepository) updateComment(comment *TeamComment) error {
	return r.db.Save(comment).Error
}

func (r *interactionRepository) deleteComment(id uuid.UUID) error {
	return r.db.Delete(&TeamComment{}, "id = ?", id).Error
}

func (r *interactionRepository) getComments(teamID uuid.UUID, limit int, offset int) ([]TeamComment, error) {
	var comments []TeamComment
	err := r.db.
		Where("team_id = ?", teamID).
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(&comments).Error
	return comments, err
}

func (r *interactionRepository) countComments(teamID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&TeamComment{}).Where("team_id = ?", teamID).Count(&count).Error
	return count, err
}
