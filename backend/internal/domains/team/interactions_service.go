package team

import (
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

/**********************
 * SERVICE INTERFACE  *
 **********************/

type teamInteractionService interface {
	LikeTeam(userID, teamID uuid.UUID) error
	UnlikeTeam(userID, teamID uuid.UUID) error
	GetTeamLikeCount(teamID uuid.UUID) (int64, error)

	ViewTeam(userID *uuid.UUID, teamID uuid.UUID) error
	GetTeamViewCount(teamID uuid.UUID) (int64, error)

	CommentTeam(userID, teamID uuid.UUID, content string, parentID *uuid.UUID) error
	GetTeamComments(teamID uuid.UUID, limit, offset int) ([]TeamComment, error)
	GetTeamCommentCount(teamID uuid.UUID) (int64, error)
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
