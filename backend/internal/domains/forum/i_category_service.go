package forum

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

/*********************************
 * CATEGORY REPOSITORY INTERFACE *
 *********************************/

type topicCategoryService interface {
	create(c *TopicCategory) error
	getByID(id uuid.UUID) (*TopicCategory, error)
	update(c *TopicCategory) error
	delete(id uuid.UUID) error
	list(limit, offset int) ([]TopicCategory, error)
}


type topicCategoryServiceImpl struct {
	repo topicCategoryRepository
	redis *redis.Client
}

func newTopicCategoryService(repo topicCategoryRepository, redis *redis.Client) topicCategoryService {
	return &topicCategoryServiceImpl{repo: repo, redis: redis}
}


// Redis Keys
func redisCategoryKey(id uuid.UUID) string {
	return fmt.Sprintf("topic_category:%s", id.String())
}

const redisCategoryListKey = "topic_category:list"

// TTL for cache (you can adjust or make it configurable)
const categoryCacheTTL = time.Hour

/**************************************
 * CATEGORY REPOSITORY IMPLEMENTATION *
 **************************************/


func (s *topicCategoryServiceImpl) create(c *TopicCategory) error {
	// Invalidate list cache
	s.redis.Del(context.Background(), redisCategoryListKey)
	return s.repo.create(c)
}

func (s *topicCategoryServiceImpl) getByID(id uuid.UUID) (*TopicCategory, error) {
	ctx := context.Background()
	key := redisCategoryKey(id)

	// Try cache
	cached, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var category TopicCategory
		if err := json.Unmarshal([]byte(cached), &category); err == nil {
			return &category, nil
		}
	}

	// Fallback to DB
	category, err := s.repo.getByID(id)
	if err != nil {
		return nil, err
	}

	// Cache result
	if data, err := json.Marshal(category); err == nil {
		s.redis.Set(ctx, key, data, categoryCacheTTL)
	}

	return category, nil
}

func (s *topicCategoryServiceImpl) update(c *TopicCategory) error {
	ctx := context.Background()
	s.redis.Del(ctx, redisCategoryKey(c.ID))
	s.redis.Del(ctx, redisCategoryListKey)
	return s.repo.update(c)
}

func (s *topicCategoryServiceImpl) delete(id uuid.UUID) error {
	ctx := context.Background()
	s.redis.Del(ctx, redisCategoryKey(id))
	s.redis.Del(ctx, redisCategoryListKey)
	return s.repo.delete(id)
}

func (s *topicCategoryServiceImpl) list(limit, offset int) ([]TopicCategory, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("%s:%d:%d", redisCategoryListKey, limit, offset)

	// Try cache
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var categories []TopicCategory
		if err := json.Unmarshal([]byte(cached), &categories); err == nil {
			return categories, nil
		}
	}

	// Fallback to DB
	categories, err := s.repo.list(limit, offset)
	if err != nil {
		return nil, err
	}

	// Cache
	if data, err := json.Marshal(categories); err == nil {
		s.redis.Set(ctx, cacheKey, data, categoryCacheTTL)
	}

	return categories, nil
}
