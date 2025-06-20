package team

import (
	"context"
	"encoding/json"
	"fmt"
	"pokemon/pkg/utils"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

/*********************
 * SERVICE INTERFACE *
 *********************/

type teamService interface {
	createTeam(team *Team) error
	getTeam(id uuid.UUID) (*Team, error)
	listTeams(userID uuid.UUID, limit int, offset int) ([]Team, error)
	updateTeam(team *Team) error
	deleteTeam(id uuid.UUID) error
}

/********************
 * REDIS KEY UTILS  *
 ********************/

func redisTeamKey(id uuid.UUID) string {
	return fmt.Sprintf("team:%s", id.String())
}

/**************************
 * SERVICE IMPLEMENTATION *
 **************************/

type service struct {
	repo  teamRepository
	redis *redis.Client
}

func newService(repo teamRepository, redis *redis.Client) teamService {
	return &service{repo: repo, redis: redis}
}

func (s *service) createTeam(team *Team) error {
	if err := team.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Persist to DB
	if err := s.repo.create(team); err != nil {
		return err
	}

	// Cache in Redis
	ctx := context.Background()
	jsonData, err := json.Marshal(team)
	if err == nil {
		s.redis.Set(ctx, redisTeamKey(team.ID), jsonData, time.Hour)
	}

	return nil
}

func (s *service) getTeam(id uuid.UUID) (*Team, error) {
	ctx := context.Background()
	key := redisTeamKey(id)

	// Try Redis
	val, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var cached Team
		if err := json.Unmarshal([]byte(val), &cached); err == nil {
			return &cached, nil
		}
	}

	// Fallback to DB
	team, err := s.repo.getByID(id)
	if err != nil {
		return nil, err
	}

	// Store in Redis
	if jsonData, err := json.Marshal(team); err == nil {
		s.redis.Set(ctx, key, jsonData, time.Hour)
	}

	return team, nil
}

func (s *service) listTeams(userID uuid.UUID, limit int, offset int) ([]Team, error) {
	// For simplicity, this skips Redis. Optional: cache with a key like `team:list:<user>:<offset>:<limit>`
	return s.repo.listByUser(userID, limit, offset)
}

func (s *service) updateTeam(team *Team) error {
	if err := team.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	err := s.repo.update(team)
	if err != nil {
		return err
	}

	// Invalidate Redis
	s.redis.Del(context.Background(), redisTeamKey(team.ID))

	return nil
}

func (s *service) deleteTeam(id uuid.UUID) error {
	err := s.repo.delete(id)
	if err != nil {
		return err
	}

	// Invalidate Redis
	s.redis.Del(context.Background(), redisTeamKey(id))

	return nil
}

/*************************
 * INTERACTIONS SERVICES *
 *************************/

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
	getSavedTeams(userID uuid.UUID, limit, offset int) ([]TeamSave, error)
	isTeamSavedByUser(userID, teamID uuid.UUID) (bool, error)

	commentTeam(userID, teamID uuid.UUID, content string, parentID *uuid.UUID) error
	getTeamComments(teamID uuid.UUID, limit, offset int) ([]TeamComment, error)
	getTeamCommentCount(teamID uuid.UUID) (int64, error)
	updateComment(comment *TeamComment) error
	deleteComment(userID, commentID uuid.UUID) error
}

/******************************
 * SERVICE IMPLEMENTATION     *
 ******************************/

type interactionService struct {
	repo  teamInteractionRepository
	redis *redis.Client
}

func newInteractionService(repo teamInteractionRepository, redis *redis.Client) teamInteractionService {
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
	utils.InvalidateCache(s.redis, utils.CacheKey("team:likes", teamID))
	return nil
}

func (s *interactionService) unlikeTeam(userID, teamID uuid.UUID) error {
	err := s.repo.deleteLike(userID, teamID)
	if err != nil {
		return err
	}
	utils.InvalidateCache(s.redis, utils.CacheKey("team:likes", teamID))
	return nil
}

func (s *interactionService) getTeamLikeCount(teamID uuid.UUID) (int64, error) {
	key := utils.CacheKey("team:likes", teamID)
	count, err := utils.GetCachedCount(s.redis, key)
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
	_ = utils.SetCachedCount(s.redis, key, count)
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
	utils.InvalidateCache(s.redis, utils.CacheKey("team:views", teamID))
	return nil
}

func (s *interactionService) getTeamViewCount(teamID uuid.UUID) (int64, error) {
	key := utils.CacheKey("team:views", teamID)
	count, err := utils.GetCachedCount(s.redis, key)
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
	_ = utils.SetCachedCount(s.redis, key, count)
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
    utils.InvalidateCache(s.redis, utils.CacheKey("team:saves", teamID))
    return nil
}

func (s *interactionService) unsaveTeam(userID, teamID uuid.UUID) error {
    err := s.repo.deleteSave(userID, teamID)
    if err != nil {
        return err
    }
    utils.InvalidateCache(s.redis, utils.CacheKey("team:saves", teamID))
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
	utils.InvalidateCache(s.redis, utils.CacheKey("team:comments", teamID))
	return nil
}

func (s *interactionService) getTeamComments(teamID uuid.UUID, limit, offset int) ([]TeamComment, error) {
	return s.repo.getComments(teamID, limit, offset)
}

func (s *interactionService) getTeamCommentCount(teamID uuid.UUID) (int64, error) {
	key := utils.CacheKey("team:comments", teamID)
	count, err := utils.GetCachedCount(s.redis, key)
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
	_ = utils.SetCachedCount(s.redis, key, count)
	return count, nil
}

func (s *interactionService) updateComment(comment *TeamComment) error {
    err := s.repo.updateComment(comment)
    if err == nil {
        utils.InvalidateCache(s.redis, utils.CacheKey("team:comments", comment.TeamID))
    }
    return err
}

func (s *interactionService) deleteComment(userID, commentID uuid.UUID) error {
    err := s.repo.deleteComment(userID, commentID)
    if err == nil {
        utils.InvalidateCache(s.redis, utils.CacheKey("team:comments", commentID))
    }
    return err
}