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

/*************************
 * INTERACTIONS SERVICES *
 *************************/

type interactionService interface {
	getInteractions(shoutID uuid.UUID) (*shoutInteractions, error)

	createLike(shoutLike *ShoutLike) error
	deleteLike(shoutID, userID uuid.UUID) error

	createComment(comment *ShoutComment) error
	updateComment(comment *ShoutComment) error
	deleteComment(commentID uuid.UUID) error

	createView(view *ShoutView) error

	createSave(save *ShoutSave) error
	deleteSave(shoutID, userID uuid.UUID) error
}

type interaction struct {
	repo  iRepository
	redis *redis.Client
}

func newInteractionService(repo iRepository, redis *redis.Client) interactionService {
	return &interaction{repo: repo, redis: redis}
}

/* Redis */
func redisShoutInteractionsKey(id uuid.UUID) string {
	return fmt.Sprintf("shout:interactions:%s", id.String())
}

func (s *interaction) getInteractions(shoutID uuid.UUID) (*shoutInteractions, error) {
	ctx := context.Background()
	key := redisShoutInteractionsKey(shoutID)

	// Try Redis
	val, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var cached shoutInteractions
		if err := json.Unmarshal([]byte(val), &cached); err == nil {
			return &cached, nil
		}
	}

	// Fallback to DB
	interactions, err := s.repo.getInteractions(shoutID)
	if err != nil {
		return nil, err
	}

	// Cache in Redis
	if jsonData, err := json.Marshal(interactions); err == nil {
		s.redis.Set(ctx, key, jsonData, time.Minute*10)
	}

	return interactions, nil
}

func (s *interaction) createLike(shoutLike *ShoutLike) error {
	err := s.repo.createLike(shoutLike)
	if err == nil {
		s.redis.Del(context.Background(), redisShoutInteractionsKey(shoutLike.ShoutID))
	}
	return err
}

func (s *interaction) deleteLike(shoutID, userID uuid.UUID) error {
	err := s.repo.deleteLike(shoutID, userID)
	if err == nil {
		s.redis.Del(context.Background(), redisShoutInteractionsKey(shoutID))
	}
	return err
}

func (s *interaction) createComment(comment *ShoutComment) error {
	err := s.repo.createComment(comment)
	if err == nil {
		s.redis.Del(context.Background(), redisShoutInteractionsKey(comment.ShoutID))
	}
	return err
}

func (s *interaction) updateComment(comment *ShoutComment) error {
	return s.repo.updateComment(comment)
}

func (s *interaction) deleteComment(commentID uuid.UUID) error {
	// You could optionally fetch the comment first to get ShoutID for invalidating Redis
	return s.repo.deleteComment(commentID)
}

func (s *interaction) createView(view *ShoutView) error {
	err := s.repo.createView(view)
	if err == nil {
		s.redis.Del(context.Background(), redisShoutInteractionsKey(view.ShoutID))
	}
	return err
}

func (s *interaction) createSave(save *ShoutSave) error {
	err := s.repo.createSave(save)
	if err == nil {
		s.redis.Del(context.Background(), redisShoutInteractionsKey(save.ShoutID))
	}
	return err
}

func (s *interaction) deleteSave(shoutID, userID uuid.UUID) error {
	err := s.repo.deleteSave(shoutID, userID)
	if err == nil {
		s.redis.Del(context.Background(), redisShoutInteractionsKey(shoutID))
	}
	return err
}