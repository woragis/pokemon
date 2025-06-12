package forum

import (
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

/**********************************
 * COMMENT LIKE SERVICE INTERFACE *
 **********************************/

type commentLikeService interface {
	create(like *CommentLike) error
	update(like *CommentLike) error
	get(commentID, userID uuid.UUID) (*CommentLike, error)
	delete(commentID, userID uuid.UUID) error
}

type commentLikeServiceStruct struct {
	repo commentLikeRepository
	redis *redis.Client
}

func newCommentLikeService(repo commentLikeRepository, redis *redis.Client) commentLikeService {
	return &commentLikeServiceStruct{repo: repo, redis: redis}
}

/***************************************
 * COMMENT LIKE SERVICE IMPLEMENTATION *
 ***************************************/

func (s *commentLikeServiceStruct) create(like *CommentLike) error {
	return s.repo.create(like)
}
func (s *commentLikeServiceStruct) update(like *CommentLike) error {
	return s.repo.update(like)
}
func (s *commentLikeServiceStruct) get(commentID, userID uuid.UUID) (*CommentLike, error) {
	return s.repo.get(commentID, userID)
}
func (s *commentLikeServiceStruct) delete(commentID, userID uuid.UUID) error {
	return s.repo.delete(commentID, userID)
}