package team

import (
	"context"
	"encoding/json"
	"fmt"
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
