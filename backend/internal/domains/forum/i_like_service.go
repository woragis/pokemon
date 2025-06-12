package forum

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

/**************************
 * LIKE SERVICE INTERFACE *
 **************************/

type topicLikeService interface {
	create(like *TopicLike) error
	update(like *TopicLike) error
	get(topicID, userID uuid.UUID) (*TopicLike, error)
	delete(topicID, userID uuid.UUID) error
}

type topicLikeServiceImpl struct {
	repo topicLikeRepository
	redis *redis.Client
}

func newTopicLikeService(repo topicLikeRepository, redis *redis.Client) topicLikeService {
	return &topicLikeServiceImpl{repo: repo, redis: redis}
}

func redisTopicLikeKey(topicID, userID uuid.UUID) string {
	return fmt.Sprintf("topiclike:%s:%s", topicID.String(), userID.String())
}

/**********************************
 * LIKE REPOSITORY IMPLEMENTATION *
 **********************************/

func (s *topicLikeServiceImpl) create(like *TopicLike) error {
	if err := s.repo.create(like); err != nil {
		return err
	}

	// Set in Redis
	key := redisTopicLikeKey(like.TopicID, like.UserID)
	data, _ := json.Marshal(like)
	s.redis.Set(context.Background(), key, data, 10*time.Minute)
	return nil
}

func (s *topicLikeServiceImpl) update(like *TopicLike) error {
	if err := s.repo.update(like); err != nil {
		return err
	}

	// Update in Redis
	key := redisTopicLikeKey(like.TopicID, like.UserID)
	data, _ := json.Marshal(like)
	s.redis.Set(context.Background(), key, data, 10*time.Minute)
	return nil
}

func (s *topicLikeServiceImpl) get(topicID, userID uuid.UUID) (*TopicLike, error) {
	ctx := context.Background()
	key := redisTopicLikeKey(topicID, userID)

	// Try cache
	cached, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var like TopicLike
		if err := json.Unmarshal([]byte(cached), &like); err == nil {
			return &like, nil
		}
	}

	// Fallback to DB
	like, err := s.repo.get(topicID, userID)
	if err != nil {
		return nil, err
	}

	// Cache result
	data, _ := json.Marshal(like)
	s.redis.Set(ctx, key, data, 10*time.Minute)

	return like, nil
}

func (s *topicLikeServiceImpl) delete(topicID, userID uuid.UUID) error {
	if err := s.repo.delete(topicID, userID); err != nil {
		return err
	}

	// Delete from Redis
	key := redisTopicLikeKey(topicID, userID)
	s.redis.Del(context.Background(), key)

	return nil
}
