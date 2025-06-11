package team

import (
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

/**********************
 * SERVICE INTERFACE  *
 **********************/

type teamInteractionService interface {
	likeTeam(userID, teamID uuid.UUID) error
	unlikeTeam(userID, teamID uuid.UUID) error
	getTeamLikeCount(teamID uuid.UUID) (int64, error)
	isTeamLikedByUser(userID, teamID uuid.UUID) (bool, error)

	viewTeam(userID *uuid.UUID, teamID uuid.UUID) error
	getTeamViewCount(teamID uuid.UUID) (int64, error)

	saveTeam(userID, teamID uuid.UUID) error
	unsaveTeam(userID, teamID uuid.UUID) error
	getSavedTeams(userID, limit, offset int) ([]TeamSave, error)
	isTeamSavedByUser(userID, teamID uuid.UUID) (bool, error)

	commentTeam(userID, teamID uuid.UUID, content string, parentID *uuid.UUID) error
	getTeamComments(teamID uuid.UUID, limit, offset int) ([]TeamComment, error)
	getTeamCommentCount(teamID uuid.UUID) (int64, error)
	updateComment(comment *TeamComment) error
	deleteComment(id uuid.UUID) error
}

/******************************
 * SERVICE IMPLEMENTATION     *
 ******************************/

type interactionService struct {
	repo  teamInteractionRepository
	redis *redis.Client
}

func newInteractionService(repo teamInteractionRepository, redis *redis.Client) *interactionService {
	return &interactionService{repo, redis}
}

/********
 * LIKE *
 ********/

func (s *interactionService) likeTeam(userID, teamID uuid.UUID) error {
	err := s.repo.createLike(&TeamLike{UserID: userID, TeamID: teamID})
	if err != nil {
		return err
	}
	s.invalidateCache(cacheKey("team:likes", teamID))
	return nil
}

func (s *interactionService) unlikeTeam(userID, teamID uuid.UUID) error {
	err := s.repo.deleteLike(userID, teamID)
	if err != nil {
		return err
	}
	s.invalidateCache(cacheKey("team:likes", teamID))
	return nil
}

func (s *interactionService) getTeamLikeCount(teamID uuid.UUID) (int64, error) {
	key := cacheKey("team:likes", teamID)
	count, err := s.getCachedCount(key)
	if err != nil {
		return 0, err
	}
	if count >= 0 {
		return count, nil
	}
	count, err = s.repo.countLikes(teamID)
	if err != nil {
		return 0, err
	}
	_ = s.setCachedCount(key, count)
	return count, nil
}

func (s *interactionService) isTeamLikedByUser(userID, teamID uuid.UUID) (bool, error) {
	return s.repo.isTeamLikedByUser(userID, teamID)
}

/********
 * VIEW *
 ********/

func (s *interactionService) viewTeam(userID *uuid.UUID, teamID uuid.UUID) error {
	view := &TeamView{
		TeamID:   teamID,
		ViewedAt: time.Now(),
	}
	if userID != nil {
		view.UserID = userID
	}
	err := s.repo.createView(view)
	if err != nil {
		return err
	}
	s.invalidateCache(cacheKey("team:views", teamID))
	return nil
}

func (s *interactionService) getTeamViewCount(teamID uuid.UUID) (int64, error) {
	key := cacheKey("team:views", teamID)
	count, err := s.getCachedCount(key)
	if err != nil {
		return 0, err
	}
	if count >= 0 {
		return count, nil
	}
	count, err = s.repo.countViews(teamID)
	if err != nil {
		return 0, err
	}
	_ = s.setCachedCount(key, count)
	return count, nil
}

/********
 * SAVE *
 ********/
 
func (s *interactionService) saveTeam(userID, teamID uuid.UUID) error {
    err := s.repo.createSave(&TeamSave{UserID: userID, TeamID: teamID})
    if err != nil {
        return err
    }
    s.invalidateCache(cacheKey("team:saves", teamID))
    return nil
}

func (s *interactionService) unsaveTeam(userID, teamID uuid.UUID) error {
    err := s.repo.deleteSave(userID, teamID)
    if err != nil {
        return err
    }
    s.invalidateCache(cacheKey("team:saves", teamID))
    return nil
}

func (s *interactionService) getSavedTeams(userID uuid.UUID, limit, offset int) ([]TeamSave, error) {
	return s.repo.listSavedTeams(userID, limit, offset)
}

func (s *interactionService) isTeamSavedByUser(userID, teamID uuid.UUID) (bool, error) {
    return s.repo.isTeamSavedByUser(userID, teamID)
}

/***********
 * COMMENT *
 ***********/

func (s *interactionService) commentTeam(userID, teamID uuid.UUID, content string, parentID *uuid.UUID) error {
	comment := &TeamComment{
		UserID:   userID,
		TeamID:   teamID,
		ParentID: parentID,
		Content:  content,
	}
	err := s.repo.createComment(comment)
	if err != nil {
		return err
	}
	s.invalidateCache(cacheKey("team:comments", teamID))
	return nil
}

func (s *interactionService) getTeamComments(teamID uuid.UUID, limit, offset int) ([]TeamComment, error) {
	return s.repo.getComments(teamID, limit, offset)
}

func (s *interactionService) getTeamCommentCount(teamID uuid.UUID) (int64, error) {
	key := cacheKey("team:comments", teamID)
	count, err := s.getCachedCount(key)
	if err != nil {
		return 0, err
	}
	if count >= 0 {
		return count, nil
	}
	count, err = s.repo.countComments(teamID)
	if err != nil {
		return 0, err
	}
	_ = s.setCachedCount(key, count)
	return count, nil
}

func (s *interactionService) updateComment(comment *TeamComment) error {
    err := s.repo.updateComment(comment)
    if err == nil {
        s.invalidateCache(cacheKey("team:comments", comment.TeamID))
    }
    return err
}

func (s *interactionService) deleteComment(id uuid.UUID) error {
    err := s.repo.deleteComment(id)
    if err == nil {
        s.invalidateCache(cacheKey("team:comments", id))
    }
    return err
}
