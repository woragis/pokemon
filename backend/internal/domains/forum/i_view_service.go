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
 * VIEW SERVICE INTERFACE *
 **************************/

type topicViewService interface {
	create(view *TopicView) error
	listByUser(userID uuid.UUID, limit, offset int) ([]TopicView, error)
}

type topicViewServiceImpl struct {
	repo topicViewRepository
	redis *redis.Client
}

func newTopicViewService(repo topicViewRepository, redis *redis.Client) topicViewService {
	return &topicViewServiceImpl{repo: repo, redis: redis}
}

func redisTopicViewKey(userID uuid.UUID, limit, offset int) string {
	return fmt.Sprintf("user:%s:topicviews:l%d:o%d", userID.String(), limit, offset)
}

/*******************************
 * VIEW SERVICE IMPLEMENTATION *
 *******************************/

func (s *topicViewServiceImpl) create(view *TopicView) error {
	// Invalidate cache after insert
	cacheKey := redisTopicViewKey(view.UserID, 10, 0) // You may want to clear multiple keys
	_ = s.redis.Del(context.Background(), cacheKey).Err()

	return s.repo.create(view)
}

func (s *topicViewServiceImpl) listByUser(userID uuid.UUID, limit, offset int) ([]TopicView, error) {
	ctx := context.Background()
	cacheKey := redisTopicViewKey(userID, limit, offset)

	// Try to get from cache
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var views []TopicView
		if err := json.Unmarshal([]byte(cached), &views); err == nil {
			return views, nil
		}
	}

	// Fallback to DB
	views, err := s.repo.listByUser(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	// Store in Redis
	if jsonData, err := json.Marshal(views); err == nil {
		_ = s.redis.Set(ctx, cacheKey, jsonData, time.Minute*5).Err()
	}

	return views, nil
}
