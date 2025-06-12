package forum

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

/*****************************
 * COMMENT SERVICE INTERFACE *
 *****************************/

type topicCommentService interface {
	create(c *TopicComment) error
	getByID(id uuid.UUID) (*TopicComment, error)
	update(c *TopicComment) error
	delete(id uuid.UUID) error
	listByTopic(topicID uuid.UUID) ([]TopicComment, error)
}

type topicCommentServiceImpl struct {
	repo topicCommentRepository
	redis *redis.Client
}

func newTopicCommentService(repo topicCommentRepository, redis *redis.Client) topicCommentService {
	return &topicCommentServiceImpl{repo: repo, redis: redis}
}

func redisTopicCommentsKey(topicID uuid.UUID) string {
	return fmt.Sprintf("topic:%s:comments", topicID)
}

/**********************************
 * COMMENT SERVICE IMPLEMENTATION *
 **********************************/

// Create clears topic comment cache
func (s *topicCommentServiceImpl) create(c *TopicComment) error {
	if err := s.repo.create(c); err != nil {
		return err
	}

	// Invalidate cached list for topic
	s.redis.Del(context.Background(), redisTopicCommentsKey(c.TopicID))
	return nil
}

func (s *topicCommentServiceImpl) getByID(id uuid.UUID) (*TopicComment, error) {
	ctx := context.Background()
	key := redisTopicCommentsKey(id)

	// Try Redis first
	val, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var comment TopicComment
		if jsonErr := json.Unmarshal([]byte(val), &comment); jsonErr == nil {
			return &comment, nil
		}
	}

	// Fallback to DB
	comment, err := s.repo.getByID(id)
	if err != nil {
		return nil, err
	}

	// Cache result
	data, _ := json.Marshal(comment)
	s.redis.Set(ctx, key, data, time.Minute*10)
	return comment, nil
}

func (s *topicCommentServiceImpl) update(c *TopicComment) error {
	if err := s.repo.update(c); err != nil {
		return err
	}

	ctx := context.Background()

	// Invalidate both single comment and topic comment list
	s.redis.Del(ctx, redisTopicCommentsKey(c.TopicID))
	return nil
}

func (s *topicCommentServiceImpl) delete(id uuid.UUID) error {
	// Get comment to find topic ID before deleting
	comment, err := s.repo.getByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.delete(id); err != nil {
		return err
	}

	ctx := context.Background()
	s.redis.Del(ctx, redisTopicCommentsKey(comment.TopicID))
	return nil
}

func (s *topicCommentServiceImpl) listByTopic(topicID uuid.UUID) ([]TopicComment, error) {
	ctx := context.Background()
	key := redisTopicCommentsKey(topicID)

	// Try Redis
	val, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var comments []TopicComment
		if jsonErr := json.Unmarshal([]byte(val), &comments); jsonErr == nil {
			return comments, nil
		}
	}

	// Fallback to DB
	comments, err := s.repo.listByTopic(topicID)
	if err != nil {
		return nil, err
	}

	// Cache result
	data, _ := json.Marshal(comments)
	s.redis.Set(ctx, key, data, time.Minute*5)
	return comments, nil
}