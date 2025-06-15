package shout

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// shoutService interface
type shoutService interface {
	createShout(shout *Shout) error
	getShout(id uuid.UUID) (*Shout, error)
	listShouts(limit, offset int) ([]Shout, error)
	listShoutsByUser(userID uuid.UUID, limit, offset int) ([]Shout, error)
	listShoutsByParent(reshoutID uuid.UUID, limit, offset int) ([]Shout, error)
	updateShout(shout *Shout) error
	deleteShout(id uuid.UUID) error
	countShoutsByUser(userID uuid.UUID) (int64, error)
}

/********************
 * REDIS KEY UTILS  *
 ********************/

func redisShoutKey(id uuid.UUID) string {
	return fmt.Sprintf("shout:%s", id.String())
}

/**************************
 * SERVICE IMPLEMENTATION *
 **************************/

type service struct {
	repo  shoutRepository
	redis *redis.Client
}

func newService(repo shoutRepository, redis *redis.Client) shoutService {
	return &service{repo: repo, redis: redis}
}

func (s *service) createShout(shout *Shout) error {
	// You can add validation here if you want (shout.Validate())

	if err := s.repo.create(shout); err != nil {
		return err
	}

	ctx := context.Background()
	jsonData, err := json.Marshal(shout)
	if err == nil {
		s.redis.Set(ctx, redisShoutKey(shout.ID), jsonData, time.Hour)
	}

	return nil
}

func (s *service) getShout(id uuid.UUID) (*Shout, error) {
	ctx := context.Background()
	key := redisShoutKey(id)

	// Try Redis cache
	val, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var cached Shout
		if err := json.Unmarshal([]byte(val), &cached); err == nil {
			return &cached, nil
		}
	}

	// Fallback to DB
	shout, err := s.repo.getById(id)
	if err != nil {
		return nil, err
	}

	// Store in Redis cache
	if jsonData, err := json.Marshal(shout); err == nil {
		s.redis.Set(ctx, key, jsonData, time.Hour)
	}

	return shout, nil
}

func (s *service) listShouts(limit, offset int) ([]Shout, error) {
	// Skip caching for list queries for simplicity
	return s.repo.list(limit, offset)
}

func (s *service) listShoutsByUser(userID uuid.UUID, limit, offset int) ([]Shout, error) {
	return s.repo.listByUser(userID, limit, offset)
}

func (s *service) listShoutsByParent(reshoutID uuid.UUID, limit, offset int) ([]Shout, error) {
	return s.repo.listByParent(reshoutID, limit, offset)
}

func (s *service) updateShout(shout *Shout) error {
	// Optional: validation here (shout.Validate())

	if err := s.repo.update(shout); err != nil {
		return err
	}

	// Invalidate Redis cache
	s.redis.Del(context.Background(), redisShoutKey(shout.ID))

	return nil
}

func (s *service) deleteShout(id uuid.UUID) error {
	if err := s.repo.delete(id); err != nil {
		return err
	}

	// Invalidate Redis cache
	s.redis.Del(context.Background(), redisShoutKey(id))

	return nil
}

func (s *service) countShoutsByUser(userID uuid.UUID) (int64, error) {
	return s.repo.countByUser(userID)
}
