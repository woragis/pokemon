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

/*************************
 * INTERACTIONS SERVICES *
 *************************/

type snapCommentService interface {
	listByUser(userID uuid.UUID, limit, offset int) ([]SnapComment, error)
	countByUser(userID uuid.UUID) (int64, error)
	create(comment *SnapComment) error
	updateStatus(id uuid.UUID, status string) error
	delete(id uuid.UUID) error
	exists(id uuid.UUID) (bool, error)
}

type snapLikeService interface {
	like(like *SnapLike) error
	unlike(snapID, userID uuid.UUID) error
	deleteAllBySnap(snapID uuid.UUID) error
	listUserLikes(userID uuid.UUID) ([]SnapLike, error)
	isLikedByUser(snapID, userID uuid.UUID) (bool, error)
}

type snapCommentLikeService interface {
	like(like *SnapCommentLike) error
	unlike(commentID, userID uuid.UUID) error
	listByComment(commentID uuid.UUID) ([]SnapCommentLike, error)
	listUserLikes(userID uuid.UUID) ([]SnapCommentLike, error)
	isLikedByUser(commentID, userID uuid.UUID) (bool, error)
}

type snapCommentServiceImpl struct {
	repo  snapCommentRepository
	redis *redis.Client
}

type snapLikeServiceImpl struct {
	repo  snapLikeRepository
	redis *redis.Client
}

type snapCommentLikeServiceImpl struct {
	repo  snapCommentLikeRepository
	redis *redis.Client
}

func newSnapCommentService(repo snapCommentRepository, redis *redis.Client) snapCommentService {
	return &snapCommentServiceImpl{repo: repo, redis: redis}
}

func newSnapLikeService(repo snapLikeRepository, redis *redis.Client) snapLikeService {
	return &snapLikeServiceImpl{repo: repo, redis: redis}
}

func newSnapCommentLikeService(repo snapCommentLikeRepository, redis *redis.Client) snapCommentLikeService {
	return &snapCommentLikeServiceImpl{repo: repo, redis: redis}
}

var SnapStatus = struct {
	Active  string
	Flagged string
	Removed string
}{
	Active:  "active",
	Flagged: "flagged",
	Removed: "removed",
}

func isValidSnapStatus(status string) bool {
	return status == SnapStatus.Active ||
		status == SnapStatus.Flagged ||
		status == SnapStatus.Removed
}

func (s *snapCommentServiceImpl) create(comment *SnapComment) error {
	if err := comment.Validate(); err != nil {
		return err
	}
	return s.repo.create(comment)
}

func (s *snapCommentServiceImpl) listByUser(userID uuid.UUID, limit, offset int) ([]SnapComment, error) {
	return s.repo.listByUser(userID, limit, offset)
}

func (s *snapCommentServiceImpl) countByUser(userID uuid.UUID) (int64, error) {
	return s.repo.countByUser(userID)
}

func (s *snapCommentServiceImpl) updateStatus(id uuid.UUID, status string) error {
	if !isValidSnapStatus(status) {
		return fmt.Errorf("invalid status: %s", status)
	}
	return s.repo.updateStatus(id, status)
}

func (s *snapCommentServiceImpl) delete(id uuid.UUID) error {
	return s.repo.delete(id)
}

func (s *snapCommentServiceImpl) exists(id uuid.UUID) (bool, error) {
	return s.repo.existsByID(id)
}

func (s *snapLikeServiceImpl) like(like *SnapLike) error {
	if err := like.Validate(); err != nil {
		return err
	}
	return s.repo.create(like)
}

func (s *snapLikeServiceImpl) unlike(snapID, userID uuid.UUID) error {
	return s.repo.delete(snapID, userID)
}

func (s *snapLikeServiceImpl) deleteAllBySnap(snapID uuid.UUID) error {
	return s.repo.deleteAllBySnap(snapID)
}

func (s *snapLikeServiceImpl) listUserLikes(userID uuid.UUID) ([]SnapLike, error) {
	return s.repo.listUserLikes(userID)
}

func (s *snapLikeServiceImpl) isLikedByUser(snapID, userID uuid.UUID) (bool, error) {
	return s.repo.exists(snapID, userID)
}

func (s *snapCommentLikeServiceImpl) like(like *SnapCommentLike) error {
	if err := like.Validate(); err != nil {
		return err
	}
	return s.repo.create(like)
}

func (s *snapCommentLikeServiceImpl) unlike(commentID, userID uuid.UUID) error {
	return s.repo.delete(commentID, userID)
}

func (s *snapCommentLikeServiceImpl) listByComment(commentID uuid.UUID) ([]SnapCommentLike, error) {
	return s.repo.listByComment(commentID)
}

func (s *snapCommentLikeServiceImpl) listUserLikes(userID uuid.UUID) ([]SnapCommentLike, error) {
	return s.repo.listUserCommentLikes(userID)
}

func (s *snapCommentLikeServiceImpl) isLikedByUser(commentID, userID uuid.UUID) (bool, error) {
	return s.repo.exists(commentID, userID)
}