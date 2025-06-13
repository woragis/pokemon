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
	get(topicID, userID uuid.UUID) (*TopicLike, error)
	delete(topicID, userID uuid.UUID) error
	count(topicID uuid.UUID, likeValue bool) (int64, error)
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

func (s *topicLikeServiceImpl) count(topicID uuid.UUID, likeValue bool) (int64, error) {
	return s.repo.count(topicID, likeValue)
}

func (s *topicLikeServiceImpl) create(like *TopicLike) error {
	if err := s.repo.create(like); err != nil {
		return err
	}

	// Cache in Redis
	key := redisTopicLikeKey(like.TopicID, like.UserID)
	data, _ := json.Marshal(like)
	_ = s.redis.Set(context.Background(), key, data, 10*time.Minute).Err()

	return nil
}

func (s *topicLikeServiceImpl) get(topicID, userID uuid.UUID) (*TopicLike, error) {
	ctx := context.Background()
	key := redisTopicLikeKey(topicID, userID)

	// Try Redis
	if cached, err := s.redis.Get(ctx, key).Result(); err == nil {
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
	if data, err := json.Marshal(like); err == nil {
		_ = s.redis.Set(ctx, key, data, 10*time.Minute).Err()
	}

	return like, nil
}

func (s *topicLikeServiceImpl) delete(topicID, userID uuid.UUID) error {
	if err := s.repo.delete(topicID, userID); err != nil {
		return err
	}

	// Invalidate Redis
	key := redisTopicLikeKey(topicID, userID)
	_ = s.redis.Del(context.Background(), key).Err()

	return nil
}