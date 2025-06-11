package blog

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

type postService interface {
	createPost(post *Post) error
	getPost(id uuid.UUID) (*Post, error)
	listPosts(userID uuid.UUID, limit int, offset int) ([]Post, error)
	updatePost(post *Post) error
	deletePost(id uuid.UUID) error
}

func redisPostKey(id uuid.UUID) string {
	return fmt.Sprintf("post:%s", id.String())
}

/**************************
 * SERVICE IMPLEMENTATION *
 **************************/

type service struct {
	repo  blogRepository
	redis *redis.Client
}

func newService(repo blogRepository, redis *redis.Client) postService {
	return &service{repo: repo, redis: redis}
}

func (s *service) createPost(post *Post) error {
	if err := post.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Persist to DB
	if err := s.repo.create(post); err != nil {
		return err
	}

	// Cache in Redis
	ctx := context.Background()
	jsonData, err := json.Marshal(post)
	if err == nil {
		s.redis.Set(ctx, redisPostKey(post.ID), jsonData, time.Hour)
	}

	return nil
}

func (s *service) getPost(id uuid.UUID) (*Post, error) {
	ctx := context.Background()
	key := redisPostKey(id)

	// Try Redis
	val, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var cached Post
		if err := json.Unmarshal([]byte(val), &cached); err == nil {
			return &cached, nil
		}
	}

	// Fallback to DB
	post, err := s.repo.getByID(id)
	if err != nil {
		return nil, err
	}

	// Store in Redis
	if jsonData, err := json.Marshal(post); err == nil {
		s.redis.Set(ctx, key, jsonData, time.Hour)
	}

	return post, nil
}

func (s *service) listPosts(userID uuid.UUID, limit int, offset int) ([]Post, error) {
	return s.repo.listByUser(userID, limit, offset)
}

func (s *service) updatePost(post *Post) error {
	if err := post.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	err := s.repo.update(post)
	if err != nil {
		return err
	}

	// Invalidate Redis
	s.redis.Del(context.Background(), redisPostKey(post.ID))

	return nil
}

func (s *service) deletePost(id uuid.UUID) error {
	err := s.repo.delete(id)
	if err != nil {
		return err
	}

	// Invalidate Redis
	s.redis.Del(context.Background(), redisPostKey(id))

	return nil
}
