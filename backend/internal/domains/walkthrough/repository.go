package walkthrough

import (
	"context"

	"gorm.io/gorm"
)

/**********************
 **********************
 ******** MAIN ********
 **********************
 **********************/

type repository interface {
	create(ctx context.Context, wt *Walkthrough) error
	getByID(ctx context.Context, id string) (*Walkthrough, error)
	list(ctx context.Context, limit, offset int) ([]Walkthrough, int64, error)
	update(ctx context.Context, wt *Walkthrough) error
	delete(ctx context.Context, id string) error

	addStep(ctx context.Context, step *WalkthroughStep) error
	updateStep(ctx context.Context, step *WalkthroughStep) error
	deleteStep(ctx context.Context, stepID string) error

	addComment(ctx context.Context, comment *WalkthroughComment) error
	listComments(ctx context.Context, walkthroughID string, limit, offset int) ([]WalkthroughComment, int64, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func newRepo(db *gorm.DB) repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) create(ctx context.Context, wt *Walkthrough) error {
	return r.db.WithContext(ctx).Create(wt).Error
}

func (r *repositoryImpl) getByID(ctx context.Context, id string) (*Walkthrough, error) {
	var wt Walkthrough
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Steps", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("step_number ASC")
		}).
		First(&wt, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &wt, nil
}

func (r *repositoryImpl) list(ctx context.Context, limit, offset int) ([]Walkthrough, int64, error) {
	var list []Walkthrough
	var count int64

	tx := r.db.WithContext(ctx).Model(&Walkthrough{})
	if err := tx.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	err := tx.Preload("User").
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (r *repositoryImpl) update(ctx context.Context, wt *Walkthrough) error {
	return r.db.WithContext(ctx).Save(wt).Error
}

func (r *repositoryImpl) delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&Walkthrough{}, "id = ?", id).Error
}

func (r *repositoryImpl) addStep(ctx context.Context, step *WalkthroughStep) error {
	return r.db.WithContext(ctx).Create(step).Error
}

func (r *repositoryImpl) updateStep(ctx context.Context, step *WalkthroughStep) error {
	return r.db.WithContext(ctx).Save(step).Error
}

func (r *repositoryImpl) deleteStep(ctx context.Context, stepID string) error {
	return r.db.WithContext(ctx).Delete(&WalkthroughStep{}, "id = ?", stepID).Error
}

func (r *repositoryImpl) addComment(ctx context.Context, comment *WalkthroughComment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

func (r *repositoryImpl) listComments(ctx context.Context, walkthroughID string, limit, offset int) ([]WalkthroughComment, int64, error) {
	var list []WalkthroughComment
	var count int64

	tx := r.db.WithContext(ctx).Model(&WalkthroughComment{}).Where("walkthrough_id = ?", walkthroughID)
	if err := tx.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	err := tx.Preload("User").
		Order("created_at ASC").
		Limit(limit).Offset(offset).
		Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

/******************************
 ******************************
 ******** INTERACTIONS ********
 ******************************
 ******************************/
