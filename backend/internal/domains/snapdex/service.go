package snapdex

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

type SnapService interface {
	create(snap *Snap) error
	update(snap *Snap) error
	delete(id uuid.UUID) error

	getByID(id uuid.UUID) (*Snap, error)
	list(limit, offset int) ([]Snap, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]Snap, error)
	countByUser(userID uuid.UUID) (int64, error)
}

type snapService struct {
	repo  snapRepository
	redis *redis.Client
}

func newSnapService(repo snapRepository, redis *redis.Client) SnapService {
	return &snapService{repo: repo, redis: redis}
}

// Cache helpers
func cacheKey(id uuid.UUID) string {
	return fmt.Sprintf("snap:%s", id.String())
}

const snapTTL = time.Minute * 10

/***************************
 * SERVICE IMPLEMENTATIONS *
 ***************************/

func (s *snapService) create(snap *Snap) error {
	if err := snap.Validate(); err != nil {
		return err
	}
	return s.repo.create(snap)
}

func (s *snapService) update(snap *Snap) error {
	if err := snap.Validate(); err != nil {
		return err
	}
	if err := s.repo.update(snap); err != nil {
		return err
	}
	_ = s.redis.Del(context.Background(), cacheKey(snap.ID)) // Invalidate cache
	return nil
}

func (s *snapService) delete(id uuid.UUID) error {
	if err := s.repo.delete(id); err != nil {
		return err
	}
	_ = s.redis.Del(context.Background(), cacheKey(id)) // Invalidate cache
	return nil
}

func (s *snapService) getByID(id uuid.UUID) (*Snap, error) {
	ctx := context.Background()
	key := cacheKey(id)

	// Try Redis cache first
	if val, err := s.redis.Get(ctx, key).Result(); err == nil {
		var snap Snap
		if err := json.Unmarshal([]byte(val), &snap); err == nil {
			return &snap, nil
		}
	}

	// Fallback to DB
	snap, err := s.repo.getById(id)
	if err != nil {
		return nil, err
	}

	// Cache result
	if data, err := json.Marshal(snap); err == nil {
		_ = s.redis.Set(ctx, key, data, snapTTL).Err()
	}

	return snap, nil
}

func (s *snapService) list(limit, offset int) ([]Snap, error) {
	return s.repo.list(limit, offset)
}

func (s *snapService) listByUser(userID uuid.UUID, limit, offset int) ([]Snap, error) {
	return s.repo.listByUser(userID, limit, offset)
}

func (s *snapService) countByUser(userID uuid.UUID) (int64, error) {
	return s.repo.countByUser(userID)
}
