package guide

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

/* GAME GUIDE */

type gameGuideService interface {
	create(guide *GameGuide) error
	getByID(id uuid.UUID) (*GameGuide, error)
	getBySlug(slug string) (*GameGuide, error)
	update(guide *GameGuide) error
	delete(id uuid.UUID) error
	list(limit, offset int) ([]GameGuide, error)
	listByAuthor(authorID uuid.UUID, limit, offset int) ([]GameGuide, error)
	countByAuthor(authorID uuid.UUID) (int64, error)
}

type gameGuideServ struct {
	repo  gameGuideRepository
	redis *redis.Client
}

func newGameGuideService(repo gameGuideRepository, redis *redis.Client) gameGuideService {
	return &gameGuideServ{repo: repo, redis: redis}
}

func (s *gameGuideServ) create(guide *GameGuide) error {
	if err := guide.Validate(); err != nil {
		return err
	}
	return s.repo.create(guide)
}

func (s *gameGuideServ) getByID(id uuid.UUID) (*GameGuide, error) {
	return s.repo.getByID(id)
}

func (s *gameGuideServ) getBySlug(slug string) (*GameGuide, error) {
	ctx := context.Background()
	cacheKey := "game_guide_slug:" + slug

	// Check Redis cache
	if cached, err := s.redis.Get(ctx, cacheKey).Result(); err == nil && cached != "" {
		var guide GameGuide
		if err := json.Unmarshal([]byte(cached), &guide); err == nil {
			return &guide, nil
		}
	}

	// Fallback to DB
	guide, err := s.repo.getBySlug(slug)
	if err != nil {
		return nil, err
	}

	// Store in Redis
	if data, err := json.Marshal(guide); err == nil {
		s.redis.Set(ctx, cacheKey, data, time.Hour)
	}

	return guide, nil
}

func (s *gameGuideServ) update(guide *GameGuide) error {
	if err := guide.Validate(); err != nil {
		return err
	}
	return s.repo.update(guide)
}

func (s *gameGuideServ) delete(id uuid.UUID) error {
	return s.repo.delete(id)
}

func (s *gameGuideServ) list(limit, offset int) ([]GameGuide, error) {
	return s.repo.list(limit, offset)
}

func (s *gameGuideServ) listByAuthor(authorID uuid.UUID, limit, offset int) ([]GameGuide, error) {
	return s.repo.listByUser(authorID, limit, offset)
}

func (s *gameGuideServ) countByAuthor(authorID uuid.UUID) (int64, error) {
	return s.repo.countByUser(authorID)
}

/* GAME GUIDE TAG */

type gameGuideTagService interface {
	create(tag *GameGuideTag) error
	getByName(name string) (*GameGuideTag, error)
	listAll() ([]GameGuideTag, error)
}

type gameGuideTagServ struct {
	repo  gameGuideTagRepository
	redis *redis.Client
}

func newGameGuideTagService(repo gameGuideTagRepository, redis *redis.Client) gameGuideTagService {
	return &gameGuideTagServ{repo: repo, redis: redis}
}

func (s *gameGuideTagServ) create(tag *GameGuideTag) error {
	if tag.Name == "" {
		return errors.New("tag name is required")
	}
	return s.repo.create(tag)
}

func (s *gameGuideTagServ) getByName(name string) (*GameGuideTag, error) {
	return s.repo.getByName(name)
}

func (s *gameGuideTagServ) listAll() ([]GameGuideTag, error) {
	ctx := context.Background()
	cacheKey := "game_guide_tags"

	if cached, err := s.redis.Get(ctx, cacheKey).Result(); err == nil && cached != "" {
		var tags []GameGuideTag
		if err := json.Unmarshal([]byte(cached), &tags); err == nil {
			return tags, nil
		}
	}

	tags, err := s.repo.listAll()
	if err != nil {
		return nil, err
	}

	if data, err := json.Marshal(tags); err == nil {
		s.redis.Set(ctx, cacheKey, data, time.Hour*2)
	}

	return tags, nil
}
