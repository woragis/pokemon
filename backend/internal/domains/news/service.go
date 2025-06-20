package news

import (
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type newsService interface {
	create(news *News) error
	get(id uuid.UUID) (*News, error)
	list(limit, offset int) ([]News, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]News, error)
	update(news *News) error
	delete(id uuid.UUID) error
	countByUser(userID uuid.UUID) (int64, error)
}

type service struct {
	repo  repository
	redis *redis.Client
}

func newService(repo repository, redis *redis.Client) newsService {
	return &service{repo: repo, redis: redis}
}

func (s *service) create(news *News) error {
	if err := news.Validate(); err != nil {
		return err
	}
	return s.repo.create(*news)
}

func (s *service) get(id uuid.UUID) (*News, error) {
	n, err := s.repo.get(id)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func (s *service) list(limit, offset int) ([]News, error) {
	return s.repo.list(limit, offset)
}

func (s *service) listByUser(userID uuid.UUID, limit, offset int) ([]News, error) {
	return s.repo.listByUser(userID, limit, offset)
}

func (s *service) update(news *News) error {
	if err := news.Validate(); err != nil {
		return err
	}
	return s.repo.update(*news)
}

func (s *service) delete(id uuid.UUID) error {
	// optionally check user ownership here if needed
	return s.repo.delete(id, uuid.Nil) // or pass userID if ownership required
}

func (s *service) countByUser(userID uuid.UUID) (int64, error) {
	return s.repo.countByUser(userID)
}
