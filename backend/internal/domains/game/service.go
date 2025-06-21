package game

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type gameService interface {
	create(ctx context.Context, game *Game) error
	getByID(ctx context.Context, id string) (*Game, error)
	list(ctx context.Context, limit, offset int) ([]Game, int64, error)
	update(ctx context.Context, game *Game) error
	delete(ctx context.Context, id string) error
}

type gameServiceImpl struct {
	db    gameRepository
	cache *redis.Client
}

func newGameService(repo gameRepository, cache *redis.Client) gameService {
	return &gameServiceImpl{db: repo, cache: cache}
}

func (s *gameServiceImpl) create(ctx context.Context, game *Game) error {
	if err := s.db.create(ctx, game); err != nil {
		return err
	}
	return nil
}

func (s *gameServiceImpl) getByID(ctx context.Context, id string) (*Game, error) {
	cacheKey := redisGameKey(id)

	// Try Redis first
	val, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var game Game
		if err := json.Unmarshal([]byte(val), &game); err == nil {
			return &game, nil
		}
	}

	// Fallback to DB
	game, err := s.db.getByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Cache result
	bytes, _ := json.Marshal(game)
	s.cache.Set(ctx, cacheKey, bytes, 10*time.Minute)

	return game, nil
}

func (s *gameServiceImpl) list(ctx context.Context, limit, offset int) ([]Game, int64, error) {
	return s.db.list(ctx, limit, offset)
}

func (s *gameServiceImpl) update(ctx context.Context, game *Game) error {
	if err := s.db.update(ctx, game); err != nil {
		return err
	}
	s.cache.Del(ctx, redisGameKey(game.ID.String()))
	return nil
}

func (s *gameServiceImpl) delete(ctx context.Context, id string) error {
	if err := s.db.delete(ctx, id); err != nil {
		return err
	}
	s.cache.Del(ctx, redisGameKey(id))
	return nil
}

func redisGameKey(id string) string {
	return fmt.Sprintf("game:%s", id)
}

type gamePokedexService interface {
	create(ctx context.Context, dex *GamePokedex) error
	getByID(ctx context.Context, id string) (*GamePokedex, error)
	getByUserAndGame(ctx context.Context, userID, gameID string) (*GamePokedex, error)
	listByUser(ctx context.Context, userID string, limit, offset int) ([]GamePokedex, int64, error)
	update(ctx context.Context, dex *GamePokedex) error
	delete(ctx context.Context, id string) error
}

type gamePokedexServiceImpl struct {
	db    gamePokedexRepository
	cache *redis.Client
}

func newGamePokedexService(repo gamePokedexRepository, cache *redis.Client) gamePokedexService {
	return &gamePokedexServiceImpl{db: repo, cache: cache}
}

func (s *gamePokedexServiceImpl) create(ctx context.Context, dex *GamePokedex) error {
	return s.db.create(ctx, dex)
}

func (s *gamePokedexServiceImpl) getByID(ctx context.Context, id string) (*GamePokedex, error) {
	cacheKey := redisDexKey(id)

	val, err := s.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var dex GamePokedex
		if err := json.Unmarshal([]byte(val), &dex); err == nil {
			return &dex, nil
		}
	}

	dex, err := s.db.getByID(ctx, id)
	if err != nil {
		return nil, err
	}

	bytes, _ := json.Marshal(dex)
	s.cache.Set(ctx, cacheKey, bytes, 10*time.Minute)

	return dex, nil
}

func (s *gamePokedexServiceImpl) getByUserAndGame(ctx context.Context, userID, gameID string) (*GamePokedex, error) {
	// Optional cache — left uncached for user-specific and possibly high-churn data
	return s.db.getByUserAndGame(ctx, userID, gameID)
}

func (s *gamePokedexServiceImpl) listByUser(ctx context.Context, userID string, limit, offset int) ([]GamePokedex, int64, error) {
	return s.db.listByUser(ctx, userID, limit, offset)
}

func (s *gamePokedexServiceImpl) update(ctx context.Context, dex *GamePokedex) error {
	if err := s.db.update(ctx, dex); err != nil {
		return err
	}
	s.cache.Del(ctx, redisDexKey(dex.ID.String()))
	return nil
}

func (s *gamePokedexServiceImpl) delete(ctx context.Context, id string) error {
	if err := s.db.delete(ctx, id); err != nil {
		return err
	}
	s.cache.Del(ctx, redisDexKey(id))
	return nil
}

func redisDexKey(id string) string {
	return fmt.Sprintf("pokedex:%s", id)
}

