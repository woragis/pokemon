package forum

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type topicService interface {
	create(topic *Topic) error
	getByID(id string) (*Topic, error)
	update(topic *Topic) error
	delete(id string) error
	listByUser(userID uuid.UUID, limit, offset int) ([]Topic, error)
	list(limit, offset int) ([]Topic, error)
}

type service struct {
	repo topicRepository
	redis *redis.Client
}

func redisTopicKey(ID uuid.UUID) string {
	return fmt.Sprintf("topic:%s", ID)
}

func newTopicService(repo topicRepository, redis *redis.Client) topicService {
	return &service{repo: repo, redis: redis}
}

func (s *service) create(topic *Topic) error {
	return s.repo.create(topic)
}

func (s *service) getByID(id string) (*Topic, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}
	return s.repo.getByID(uid)
}

func (s *service) update(topic *Topic) error {
	if err := topic.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	err := s.repo.update(topic)
	if err != nil {
		return err
	}

	// Invalidate Redis
	s.redis.Del(context.Background(), redisTopicKey(topic.ID))

	return nil
}

func (s *service) delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid UUID: %w", err)
	}
	return s.repo.delete(uid)
}

func (s *service) listByUser(userID uuid.UUID, limit, offset int) ([]Topic, error) {
	return s.repo.listByUser(userID, limit, offset)
}

func (s *service) list(limit, offset int) ([]Topic, error) {
	return s.repo.list(limit, offset)
}
