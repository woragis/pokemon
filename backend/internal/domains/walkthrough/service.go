package walkthrough

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type service interface {
	createWalkthrough(ctx context.Context, wt *Walkthrough) error
	getWalkthrough(ctx context.Context, id string) (*Walkthrough, error)
	listWalkthroughs(ctx context.Context, limit, offset int) ([]Walkthrough, int64, error)
	updateWalkthrough(ctx context.Context, wt *Walkthrough) error
	deleteWalkthrough(ctx context.Context, id string) error

	addStep(ctx context.Context, step *WalkthroughStep) error
	updateStep(ctx context.Context, step *WalkthroughStep) error
	deleteStep(ctx context.Context, stepID string) error

	addComment(ctx context.Context, comment *WalkthroughComment) error
	listComments(ctx context.Context, walkthroughID string, limit, offset int) ([]WalkthroughComment, int64, error)
}

type serviceImpl struct {
	repo repository
	redis *redis.Client
}

func newServ(repo repository, redis *redis.Client) service {
	return &serviceImpl{repo: repo, redis: redis}
}

func (s *serviceImpl) createWalkthrough(ctx context.Context, wt *Walkthrough) error {
	if wt.Title == "" || wt.Game == "" {
		return errors.New("title and game are required")
	}
	return s.repo.create(ctx, wt)
}

func (s *serviceImpl) getWalkthrough(ctx context.Context, id string) (*Walkthrough, error) {
	return s.repo.getByID(ctx, id)
}

func (s *serviceImpl) listWalkthroughs(ctx context.Context, limit, offset int) ([]Walkthrough, int64, error) {
	return s.repo.list(ctx, limit, offset)
}

func (s *serviceImpl) updateWalkthrough(ctx context.Context, wt *Walkthrough) error {
	if wt.ID == uuid.Nil {
		return errors.New("invalid walkthrough ID")
	}
	return s.repo.update(ctx, wt)
}

func (s *serviceImpl) deleteWalkthrough(ctx context.Context, id string) error {
	return s.repo.delete(ctx, id)
}

func (s *serviceImpl) addStep(ctx context.Context, step *WalkthroughStep) error {
	if step.WalkthroughID == uuid.Nil || step.Title == "" || step.Content == "" {
		return errors.New("invalid walkthrough step")
	}
	return s.repo.addStep(ctx, step)
}

func (s *serviceImpl) updateStep(ctx context.Context, step *WalkthroughStep) error {
	if step.ID == uuid.Nil {
		return errors.New("invalid step ID")
	}
	return s.repo.updateStep(ctx, step)
}

func (s *serviceImpl) deleteStep(ctx context.Context, stepID string) error {
	if stepID == "" {
		return errors.New("step ID required")
	}
	return s.repo.deleteStep(ctx, stepID)
}

func (s *serviceImpl) addComment(ctx context.Context, comment *WalkthroughComment) error {
	if comment.WalkthroughID == uuid.Nil || comment.UserID == uuid.Nil || comment.Content == "" {
		return errors.New("invalid comment")
	}
	return s.repo.addComment(ctx, comment)
}

func (s *serviceImpl) listComments(ctx context.Context, walkthroughID string, limit, offset int) ([]WalkthroughComment, int64, error) {
	if walkthroughID == "" {
		return nil, 0, errors.New("walkthrough ID is required")
	}
	return s.repo.listComments(ctx, walkthroughID, limit, offset)
}
