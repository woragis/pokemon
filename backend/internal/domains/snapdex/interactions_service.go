package snapdex

import (
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type interactionServ interface {
	// Commenting
	listSnapCommentsByUser(userID uuid.UUID, limit, offset int) ([]SnapComment, error)
	countSnapCommentsByUser(userID uuid.UUID) (int64, error)   // Count
	createSnapComment(comment *SnapComment) error              // Create
	updateSnapCommentStatus(id uuid.UUID, status string) error // Update
	deleteSnapComment(id uuid.UUID) error                      // Delete
	snapCommentExists(id uuid.UUID) (bool, error)              // Exists

	// Snap likes
	likeSnap(like *SnapLike) error                            // Create
	unlikeSnap(snapID, userID uuid.UUID) error                // Delete
	deleteAllLikesBySnap(snapID uuid.UUID) error              // Delete
	userSnapLikes(userID uuid.UUID) ([]SnapLike, error)       // Exists
	isSnapLikedByUser(snapID, userID uuid.UUID) (bool, error) // Exists

	// Comment likes
	likeSnapComment(like *SnapCommentLike) error                    // Create
	unlikeSnapComment(commentID, userID uuid.UUID) error            // Delete
	commentLikes(commentID uuid.UUID) ([]SnapCommentLike, error)    // Exists
	isCommentLikedByUser(commentID, userID uuid.UUID) (bool, error) // Exists
}

type interactionService struct {
	snapCommentRepo      snapCommentRepository
	snapLikeRepo         snapLikeRepository
	snapCommentLikeRepo  snapCommentLikeRepository
	redis                *redis.Client
}

func newIServ(
	snapCommentRepo snapCommentRepository,
	snapLikeRepo snapLikeRepository,
	snapCommentLikeRepo snapCommentLikeRepository,
	redis *redis.Client,
) interactionServ {
	return &interactionService{
		snapCommentRepo:      snapCommentRepo,
		snapLikeRepo:         snapLikeRepo,
		snapCommentLikeRepo:  snapCommentLikeRepo,
		redis:                redis,
	}
}

func (s *interactionService) createSnapComment(comment *SnapComment) error {
	if err := comment.Validate(); err != nil {
		return err
	}
	return s.snapCommentRepo.create(comment)
}

func (s *interactionService) listSnapCommentsByUser(userID uuid.UUID, limit, offset int) ([]SnapComment, error) {
	return s.snapCommentRepo.listByUser(userID, limit, offset)
}

func (s *interactionService) countSnapCommentsByUser(userID uuid.UUID) (int64, error) {
	return s.snapCommentRepo.countByUser(userID)
}

var SnapStatus = struct {
	Active string
	Flagged string
	Removed string
}{
	Active: "active",
	Flagged: "flagged",
	Removed: "removed",
}

func isValidSnapStatus(status string) bool {
	return status == SnapStatus.Active ||
		status == SnapStatus.Flagged ||
		status == SnapStatus.Removed
}

func (s *interactionService) updateSnapCommentStatus(id uuid.UUID, status string) error {
	return s.snapCommentRepo.updateStatus(id, status)
}

func (s *interactionService) deleteSnapComment(id uuid.UUID) error {
	return s.snapCommentRepo.delete(id)
}

func (s *interactionService) snapCommentExists(id uuid.UUID) (bool, error) {
	return s.snapCommentRepo.existsByID(id)
}

func (s *interactionService) likeSnap(like *SnapLike) error {
	if err := like.Validate(); err != nil {
		return err
	}
	return s.snapLikeRepo.create(like)
}

func (s *interactionService) unlikeSnap(snapID, userID uuid.UUID) error {
	return s.snapLikeRepo.delete(snapID, userID)
}

func (s *interactionService) deleteAllLikesBySnap(snapID uuid.UUID) error {
	return s.snapLikeRepo.deleteAllBySnap(snapID)
}

func (s *interactionService) userSnapLikes(userID uuid.UUID) ([]SnapLike, error) {
	return s.snapLikeRepo.listUserLikes(userID)
}

func (s *interactionService) isSnapLikedByUser(snapID, userID uuid.UUID) (bool, error) {
	return s.snapLikeRepo.exists(snapID, userID)
}

func (s *interactionService) likeSnapComment(like *SnapCommentLike) error {
	if err := like.Validate(); err != nil {
		return err
	}
	return s.snapCommentLikeRepo.create(like)
}

func (s *interactionService) unlikeSnapComment(commentID, userID uuid.UUID) error {
	return s.snapCommentLikeRepo.delete(commentID, userID)
}

func (s *interactionService) commentLikes(commentID uuid.UUID) ([]SnapCommentLike, error) {
	return s.snapCommentLikeRepo.listByComment(commentID)
}

func (s *interactionService) userCommentLikes(userID uuid.UUID) ([]SnapCommentLike, error) {
	return s.snapCommentLikeRepo.listUserCommentLikes(userID)
}

func (s *interactionService) isCommentLikedByUser(commentID, userID uuid.UUID) (bool, error) {
	return s.snapCommentLikeRepo.exists(commentID, userID)
}
