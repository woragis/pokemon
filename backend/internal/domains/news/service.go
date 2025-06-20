package news

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type service interface {
	create(news *News) error
	get(id uuid.UUID) (*News, error)
	list(limit, offset int) ([]News, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]News, error)
	update(news *News) error
	delete(id uuid.UUID) error
	countByUser(userID uuid.UUID) (int64, error)
}

type serviceImpl struct {
	repo  repository
	redis *redis.Client
}

func newServ(repo repository, redis *redis.Client) service {
	return &serviceImpl{repo: repo, redis: redis}
}

const cacheTTL = time.Minute * 5

func (s *serviceImpl) create(news *News) error {
	if err := news.Validate(); err != nil {
		return err
	}

	if err := s.repo.create(*news); err != nil {
		return err
	}

	ctx := context.Background()
	s.redis.Del(ctx,
		fmt.Sprintf("news:list"),
		fmt.Sprintf("news:user:%s", news.UserID),
		fmt.Sprintf("news:count:user:%s", news.UserID),
	)

	return nil
}

func (s *serviceImpl) get(id uuid.UUID) (*News, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("news:get:%s", id.String())

	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var n News
		if json.Unmarshal([]byte(cached), &n) == nil {
			return &n, nil
		}
	}

	n, err := s.repo.get(id)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(n)
	s.redis.Set(ctx, cacheKey, data, cacheTTL)

	return &n, nil
}

func (s *serviceImpl) list(limit, offset int) ([]News, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("news:list:%d:%d", limit, offset)

	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var result []News
		if json.Unmarshal([]byte(cached), &result) == nil {
			return result, nil
		}
	}

	result, err := s.repo.list(limit, offset)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(result)
	s.redis.Set(ctx, cacheKey, data, cacheTTL)

	return result, nil
}

func (s *serviceImpl) listByUser(userID uuid.UUID, limit, offset int) ([]News, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("news:user:%s:%d:%d", userID, limit, offset)

	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var result []News
		if json.Unmarshal([]byte(cached), &result) == nil {
			return result, nil
		}
	}

	result, err := s.repo.listByUser(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(result)
	s.redis.Set(ctx, cacheKey, data, cacheTTL)

	return result, nil
}

func (s *serviceImpl) update(news *News) error {
	if err := news.Validate(); err != nil {
		return err
	}

	if err := s.repo.update(*news); err != nil {
		return err
	}

	ctx := context.Background()
	s.redis.Del(ctx,
		fmt.Sprintf("news:get:%s", news.ID),
		fmt.Sprintf("news:list"),
		fmt.Sprintf("news:user:%s", news.UserID),
	)

	return nil
}

func (s *serviceImpl) delete(id uuid.UUID) error {
	n, err := s.repo.get(id)
	if err != nil {
		return err
	}

	if err := s.repo.delete(id, n.UserID); err != nil {
		return err
	}

	ctx := context.Background()
	s.redis.Del(ctx,
		fmt.Sprintf("news:get:%s", id),
		fmt.Sprintf("news:list"),
		fmt.Sprintf("news:user:%s", n.UserID),
		fmt.Sprintf("news:count:user:%s", n.UserID),
	)

	return nil
}

func (s *serviceImpl) countByUser(userID uuid.UUID) (int64, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("news:count:user:%s", userID)

	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var count int64
		if json.Unmarshal([]byte(cached), &count) == nil {
			return count, nil
		}
	}

	count, err := s.repo.countByUser(userID)
	if err != nil {
		return 0, err
	}

	data, _ := json.Marshal(count)
	s.redis.Set(ctx, cacheKey, data, cacheTTL)

	return count, nil
}
