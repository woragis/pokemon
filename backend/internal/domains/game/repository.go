package game

import (
	"context"

	"gorm.io/gorm"
)

type GameRepository interface {
	Create(ctx context.Context, game *Game) error
	GetByID(ctx context.Context, id string) (*Game, error)
	List(ctx context.Context, limit, offset int) ([]Game, int64, error)
	Update(ctx context.Context, game *Game) error
	Delete(ctx context.Context, id string) error
}

type GameRepo struct {
	db *gorm.DB
}

func NewGameRepo(db *gorm.DB) GameRepository {
	return &GameRepo{db: db}
}

func (r *GameRepo) Create(ctx context.Context, game *Game) error {
	return r.db.WithContext(ctx).Create(game).Error
}

func (r *GameRepo) GetByID(ctx context.Context, id string) (*Game, error) {
	var game Game
	if err := r.db.WithContext(ctx).First(&game, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &game, nil
}

func (r *GameRepo) List(ctx context.Context, limit, offset int) ([]Game, int64, error) {
	var games []Game
	var count int64

	tx := r.db.WithContext(ctx).Model(&Game{})
	if err := tx.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := tx.Order("released_at DESC").Limit(limit).Offset(offset).Find(&games).Error; err != nil {
		return nil, 0, err
	}

	return games, count, nil
}

func (r *GameRepo) Update(ctx context.Context, game *Game) error {
	return r.db.WithContext(ctx).Save(game).Error
}

func (r *GameRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&Game{}, "id = ?", id).Error
}

/****************
 * GAME POKEDEX *
 ****************/

type GamePokedexRepository interface {
	Create(ctx context.Context, pokedex *GamePokedex) error
	GetByID(ctx context.Context, id string) (*GamePokedex, error)
	GetByUserAndGame(ctx context.Context, userID, gameID string) (*GamePokedex, error)
	ListByUser(ctx context.Context, userID string, limit, offset int) ([]GamePokedex, int64, error)
	Update(ctx context.Context, pokedex *GamePokedex) error
	Delete(ctx context.Context, id string) error
}

type GamePokedexRepo struct {
	db *gorm.DB
}

func NewGamePokedexRepo(db *gorm.DB) GamePokedexRepository {
	return &GamePokedexRepo{db: db}
}

func (r *GamePokedexRepo) Create(ctx context.Context, pokedex *GamePokedex) error {
	return r.db.WithContext(ctx).Create(pokedex).Error
}

func (r *GamePokedexRepo) GetByID(ctx context.Context, id string) (*GamePokedex, error) {
	var pdx GamePokedex
	if err := r.db.WithContext(ctx).Preload("Game").Preload("User").First(&pdx, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &pdx, nil
}

func (r *GamePokedexRepo) GetByUserAndGame(ctx context.Context, userID, gameID string) (*GamePokedex, error) {
	var pdx GamePokedex
	if err := r.db.WithContext(ctx).Preload("Game").Preload("User").Where("user_id = ? AND game_id = ?", userID, gameID).First(&pdx).Error; err != nil {
		return nil, err
	}
	return &pdx, nil
}

func (r *GamePokedexRepo) ListByUser(ctx context.Context, userID string, limit, offset int) ([]GamePokedex, int64, error) {
	var entries []GamePokedex
	var count int64

	tx := r.db.WithContext(ctx).Model(&GamePokedex{}).Where("user_id = ?", userID)
	if err := tx.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := tx.Preload("Game").Preload("User").Order("created_at DESC").Limit(limit).Offset(offset).Find(&entries).Error; err != nil {
		return nil, 0, err
	}

	return entries, count, nil
}

func (r *GamePokedexRepo) Update(ctx context.Context, pokedex *GamePokedex) error {
	return r.db.WithContext(ctx).Save(pokedex).Error
}

func (r *GamePokedexRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&GamePokedex{}, "id = ?", id).Error
}
