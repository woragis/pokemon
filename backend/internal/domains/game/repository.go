package game

import (
	"context"

	"gorm.io/gorm"
)

/**********************
 **********************
 ******** MAIN ********
 **********************
 **********************/

type gameRepository interface {
	list(ctx context.Context, limit, offset int) ([]Game, int64, error)

	create(ctx context.Context, game *Game) error
	getByID(ctx context.Context, id string) (*Game, error)
	update(ctx context.Context, game *Game) error
	delete(ctx context.Context, id string) error
}

type gameRepoImpl struct {
	db *gorm.DB
}

func newGameRepo(db *gorm.DB) gameRepository {
	return &gameRepoImpl{db: db}
}

func (r *gameRepoImpl) create(ctx context.Context, game *Game) error {
	return r.db.WithContext(ctx).Create(game).Error
}

func (r *gameRepoImpl) getByID(ctx context.Context, id string) (*Game, error) {
	var game Game
	if err := r.db.WithContext(ctx).First(&game, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &game, nil
}

func (r *gameRepoImpl) list(ctx context.Context, limit, offset int) ([]Game, int64, error) {
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

func (r *gameRepoImpl) update(ctx context.Context, game *Game) error {
	return r.db.WithContext(ctx).Save(game).Error
}

func (r *gameRepoImpl) delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&Game{}, "id = ?", id).Error
}

/****************
 * GAME POKEDEX *
 ****************/

type gamePokedexRepository interface {
	create(ctx context.Context, pokedex *GamePokedex) error
	getByID(ctx context.Context, id string) (*GamePokedex, error)
	getByUserAndGame(ctx context.Context, userID, gameID string) (*GamePokedex, error)
	listByUser(ctx context.Context, userID string, limit, offset int) ([]GamePokedex, int64, error)
	update(ctx context.Context, pokedex *GamePokedex) error
	delete(ctx context.Context, id string) error
}

type gamePokedexRepoImpl struct {
	db *gorm.DB
}

func newGamePokedexRepo(db *gorm.DB) gamePokedexRepository {
	return &gamePokedexRepoImpl{db: db}
}

func (r *gamePokedexRepoImpl) create(ctx context.Context, pokedex *GamePokedex) error {
	return r.db.WithContext(ctx).Create(pokedex).Error
}

func (r *gamePokedexRepoImpl) getByID(ctx context.Context, id string) (*GamePokedex, error) {
	var pdx GamePokedex
	if err := r.db.WithContext(ctx).Preload("Game").Preload("User").First(&pdx, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &pdx, nil
}

func (r *gamePokedexRepoImpl) getByUserAndGame(ctx context.Context, userID, gameID string) (*GamePokedex, error) {
	var pdx GamePokedex
	if err := r.db.WithContext(ctx).Preload("Game").Preload("User").Where("user_id = ? AND game_id = ?", userID, gameID).First(&pdx).Error; err != nil {
		return nil, err
	}
	return &pdx, nil
}

func (r *gamePokedexRepoImpl) listByUser(ctx context.Context, userID string, limit, offset int) ([]GamePokedex, int64, error) {
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

func (r *gamePokedexRepoImpl) update(ctx context.Context, pokedex *GamePokedex) error {
	return r.db.WithContext(ctx).Save(pokedex).Error
}

func (r *gamePokedexRepoImpl) delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&GamePokedex{}, "id = ?", id).Error
}

/******************************
 ******************************
 ******** INTERACTIONS ********
 ******************************
 ******************************/
