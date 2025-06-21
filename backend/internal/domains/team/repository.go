package team

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/**********************
 **********************
 ******** MAIN ********
 **********************
 **********************/

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

/************************
 * REPOSITORY INTERFACE *
 ************************/

type teamInteractionRepository interface {
	// Likes
	createLike(like *TeamLike) error
	deleteLike(userID, teamID uuid.UUID) error
	countLikes(teamID uuid.UUID) (int64, error)
	isTeamLikedByUser(userID, teamID uuid.UUID) (bool, error)

	// Views
	createView(view *TeamView) error
	countViews(teamID uuid.UUID) (int64, error)

	// Saves
	createSave(save *TeamSave) error
	deleteSave(userID, teamID uuid.UUID) error
	listSavedTeams(userID uuid.UUID, limit, offset int) ([]TeamSave, error)
	isTeamSavedByUser(userID, teamID uuid.UUID) (bool, error)

	// Comments
	createComment(comment *TeamComment) error
	updateComment(comment *TeamComment) error
	deleteComment(userID, commentID uuid.UUID) error
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

func (r *interactionRepository) isTeamLikedByUser(userID, teamID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&TeamLike{}).
		Where("user_id = ? AND team_id = ?", userID, teamID).
		Count(&count).Error
	return count > 0, err
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

func (r *interactionRepository) listSavedTeams(userID uuid.UUID, limit, offset int) ([]TeamSave, error) {
	var saves []TeamSave
	err := r.db.
		Where("user_id = ?", userID).
		Preload("Team").
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(&saves).Error
	return saves, err
}

func (r *interactionRepository) isTeamSavedByUser(userID, teamID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&TeamSave{}).
		Where("user_id = ? AND team_id = ?", userID, teamID).
		Count(&count).Error
	return count > 0, err
}

// --- Comments ---
func (r *interactionRepository) createComment(comment *TeamComment) error {
	return r.db.Create(comment).Error
}

func (r *interactionRepository) updateComment(comment *TeamComment) error {
	return r.db.Save(comment).Error
}

func (r *interactionRepository) deleteComment(userID, commentID uuid.UUID) error {
	return r.db.Delete(&TeamComment{}, "id = ? AND user_id = ?", userID, commentID).Error
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

/******************************
 ******************************
 ******** INTERACTIONS ********
 ******************************
 ******************************/
