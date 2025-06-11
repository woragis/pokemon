package forum

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type topicService interface {
	create(topic *Topic) error
	getByID(id string) (*Topic, error)
	update(id string, updates map[string]interface{}) error
	delete(id string) error
	listByUser(userID uuid.UUID, limit, offset int) ([]Topic, error)
	list(limit, offset int) ([]Topic, error)
}

type service struct {
	repo topicRepository
	redis *redis.Client
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

func (s *service) update(id string, updates map[string]interface{}) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid UUID: %w", err)
	}

	topic, err := s.repo.getByID(uid)
	if err != nil {
		return err
	}

	// apply updates to topic model
	for key, value := range updates {
		switch key {
		case "title":
			if v, ok := value.(string); ok {
				topic.Title = v
			}
		case "content":
			if v, ok := value.(string); ok {
				topic.Content = v
			}
		// Add more fields here as needed
		}
	}

	return s.repo.update(topic)
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
