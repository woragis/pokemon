package favoritepokemon

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type favoritepokemonService interface {
	create(mon *FavoritePokemon) error
	listByPopular() ([]FavoritePokemon, error)
	listByUser(userID uuid.UUID) ([]FavoritePokemon, error)
	getByID(id uuid.UUID) (*FavoritePokemon, error)
	update(mon *FavoritePokemon) error
	delete(id uuid.UUID) error
}

type service struct {
	repo  favoritepokemonRepository
	redis *redis.Client
}

func newFavoritePokemonService(repo favoritepokemonRepository, redis *redis.Client) favoritepokemonService {
	return &service{
		repo:  repo,
		redis: redis,
	}
}

func (s *service) getCacheKey(id uuid.UUID) string {
	return fmt.Sprintf("favoritepokemon:%s", id.String())
}

func (s *service) cacheFavorite(mon *FavoritePokemon) {
	data, _ := json.Marshal(mon)
	s.redis.Set(context.Background(), s.getCacheKey(mon.ID), data, 15*time.Minute)
}

func (s *service) deleteCache(id uuid.UUID) {
	s.redis.Del(context.Background(), s.getCacheKey(id))
}

func (s *service) create(mon *FavoritePokemon) error {
	if err := mon.Validate(); err != nil {
		return err
	}
	err := s.repo.create(mon)
	if err == nil {
		s.cacheFavorite(mon)
	}
	return err
}

func (s *service) listByPopular() ([]FavoritePokemon, error) {
	return s.repo.listByPopular()
}

func (s *service) listByUser(userID uuid.UUID) ([]FavoritePokemon, error) {
	return s.repo.listByUser(userID)
}

func (s *service) getByID(id uuid.UUID) (*FavoritePokemon, error) {
	ctx := context.Background()
	cacheKey := s.getCacheKey(id)

	// Try cache
	if data, err := s.redis.Get(ctx, cacheKey).Result(); err == nil {
		var mon FavoritePokemon
		if err := json.Unmarshal([]byte(data), &mon); err == nil {
			return &mon, nil
		}
	}

	// Fallback to DB
	mon, err := s.repo.getByID(id)
	if err == nil {
		s.cacheFavorite(mon)
	}
	return mon, err
}

func (s *service) update(mon *FavoritePokemon) error {
	if err := mon.Validate(); err != nil {
		return err
	}
	err := s.repo.update(mon)
	if err == nil {
		s.cacheFavorite(mon)
	}
	return err
}

func (s *service) delete(id uuid.UUID) error {
	err := s.repo.delete(id)
	if err == nil {
		s.deleteCache(id)
	}
	return err
}
