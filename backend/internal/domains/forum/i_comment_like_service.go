package forum

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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
	count(commentID uuid.UUID) (likes, dislikes int64, err error)
}

type commentLikeServiceImpl struct {
	repo commentLikeRepository
	redis *redis.Client
}

func newCommentLikeService(repo commentLikeRepository, redis *redis.Client) commentLikeService {
	return &commentLikeServiceImpl{repo: repo, redis: redis}
}

const commentLikeTTL = time.Hour * 1 // cache TTL

func redisCommentLikeKey(commentID, userID uuid.UUID) string {
	return fmt.Sprintf("comment_like:%s:%s", commentID, userID)
}

/***************************************
 * COMMENT LIKE SERVICE IMPLEMENTATION *
 ***************************************/

func (s *commentLikeServiceImpl) create(like *CommentLike) error {
	if err := s.repo.create(like); err != nil {
		return err
	}
	// Cache it
	key := redisCommentLikeKey(like.CommentID, like.UserID)
	data, _ := json.Marshal(like)
	s.redis.Set(context.Background(), key, data, commentLikeTTL)
	return nil
}

func (s *commentLikeServiceImpl) update(like *CommentLike) error {
	if err := s.repo.update(like); err != nil {
		return err
	}
	// Update cache
	key := redisCommentLikeKey(like.CommentID, like.UserID)
	data, _ := json.Marshal(like)
	s.redis.Set(context.Background(), key, data, commentLikeTTL)
	return nil
}

func (s *commentLikeServiceImpl) get(commentID, userID uuid.UUID) (*CommentLike, error) {
	ctx := context.Background()
	key := redisCommentLikeKey(commentID, userID)

	// Try cache
	cached, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var like CommentLike
		if jsonErr := json.Unmarshal([]byte(cached), &like); jsonErr == nil {
			return &like, nil
		}
	}

	// Fallback to DB
	like, err := s.repo.get(commentID, userID)
	if err != nil || like == nil {
		return like, err
	}

	// Store in Redis
	data, _ := json.Marshal(like)
	s.redis.Set(ctx, key, data, commentLikeTTL)
	return like, nil
}

func (s *commentLikeServiceImpl) count(commentID uuid.UUID) (int64, int64, error) {
	return s.repo.count(commentID)
}

func (s *commentLikeServiceImpl) delete(commentID, userID uuid.UUID) error {
	if err := s.repo.delete(commentID, userID); err != nil {
		return err
	}
	key := redisCommentLikeKey(commentID, userID)
	s.redis.Del(context.Background(), key)
	return nil
}
