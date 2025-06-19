package forum

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type topicService interface {
	create(topic *Topic) error
	getByID(id string) (*Topic, error)
	update(topic *Topic) error
	delete(id string) error
	listByUser(userID uuid.UUID, limit, offset int) ([]Topic, error)
	list(limit, offset int) ([]Topic, error)
}

type service struct {
	repo topicRepository
	redis *redis.Client
}

func redisTopicKey(ID uuid.UUID) string {
	return fmt.Sprintf("topic:%s", ID)
}

func newTopicService(repo topicRepository, redis *redis.Client) topicService {
	return &service{repo: repo, redis: redis}
}

func (s *service) create(topic *Topic) error {
	return s.repo.create(topic)
}

func (s *service) getByID(id string) (*Topic, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}
	return s.repo.getByID(uid)
}

func (s *service) update(topic *Topic) error {
	if err := topic.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	err := s.repo.update(topic)
	if err != nil {
		return err
	}

	// Invalidate Redis
	s.redis.Del(context.Background(), redisTopicKey(topic.ID))

	return nil
}

func (s *service) delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid UUID: %w", err)
	}
	return s.repo.delete(uid)
}

func (s *service) listByUser(userID uuid.UUID, limit, offset int) ([]Topic, error) {
	return s.repo.listByUser(userID, limit, offset)
}

func (s *service) list(limit, offset int) ([]Topic, error) {
	return s.repo.list(limit, offset)
}

/*************************
 * INTERACTIONS SERVICES *
 *************************/

/*****************************
 * COMMENT SERVICE INTERFACE *
 *****************************/

type topicCommentService interface {
	create(c *TopicComment) error
	getByID(id uuid.UUID) (*TopicComment, error)
	update(c *TopicComment) error
	delete(id uuid.UUID) error
	listByTopic(topicID uuid.UUID) ([]TopicComment, error)
	countByTopic(topicID uuid.UUID) (int64, error)
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

func redisTopicCommentKey(commentID uuid.UUID) string {
	return fmt.Sprintf("forum:comment:%s", commentID)
}

func (s *topicCommentServiceImpl) getByID(id uuid.UUID) (*TopicComment, error) {
	ctx := context.Background()
	key := redisTopicCommentKey(id)

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
	s.redis.Del(ctx, redisTopicCommentKey(c.ID))
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
	s.redis.Del(ctx, redisTopicCommentKey(id))
	return nil
}

func redisTopicCommentsCountKey(topicID uuid.UUID) string {
	return fmt.Sprintf("forum:topic:%s:commentscount", topicID)
}
func (s *topicCommentServiceImpl) countByTopic(topicID uuid.UUID) (int64, error) {
	ctx := context.Background()
	key := redisTopicCommentsCountKey(topicID)

	val, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var count int64
		if jsonErr := json.Unmarshal([]byte(val), &count); jsonErr == nil {
			return count, nil
		}
	}

	// Fallback to DB
	count, err := s.repo.count(topicID)
	if err != nil {
		return 0, err
	}

	// Cache result
	s.redis.Set(ctx, key, count, time.Minute*5)
	return count, nil
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

/*********************************
 * CATEGORY REPOSITORY INTERFACE *
 *********************************/

type topicCategoryService interface {
	create(c *TopicCategory) error
	getByID(id uuid.UUID) (*TopicCategory, error)
	update(c *TopicCategory) error
	delete(id uuid.UUID) error
	list(limit, offset int) ([]TopicCategory, error)
}


type topicCategoryServiceImpl struct {
	repo topicCategoryRepository
	redis *redis.Client
}

func newTopicCategoryService(repo topicCategoryRepository, redis *redis.Client) topicCategoryService {
	return &topicCategoryServiceImpl{repo: repo, redis: redis}
}


// Redis Keys
func redisCategoryKey(id uuid.UUID) string {
	return fmt.Sprintf("topic_category:%s", id.String())
}

const redisCategoryListKey = "topic_category:list"

// TTL for cache (you can adjust or make it configurable)
const categoryCacheTTL = time.Hour

/**************************************
 * CATEGORY REPOSITORY IMPLEMENTATION *
 **************************************/


func (s *topicCategoryServiceImpl) create(c *TopicCategory) error {
	// Invalidate list cache
	s.redis.Del(context.Background(), redisCategoryListKey)
	return s.repo.create(c)
}

func (s *topicCategoryServiceImpl) getByID(id uuid.UUID) (*TopicCategory, error) {
	ctx := context.Background()
	key := redisCategoryKey(id)

	// Try cache
	cached, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var category TopicCategory
		if err := json.Unmarshal([]byte(cached), &category); err == nil {
			return &category, nil
		}
	}

	// Fallback to DB
	category, err := s.repo.getByID(id)
	if err != nil {
		return nil, err
	}

	// Cache result
	if data, err := json.Marshal(category); err == nil {
		s.redis.Set(ctx, key, data, categoryCacheTTL)
	}

	return category, nil
}

func (s *topicCategoryServiceImpl) update(c *TopicCategory) error {
	ctx := context.Background()
	s.redis.Del(ctx, redisCategoryKey(c.ID))
	s.redis.Del(ctx, redisCategoryListKey)
	return s.repo.update(c)
}

func (s *topicCategoryServiceImpl) delete(id uuid.UUID) error {
	ctx := context.Background()
	s.redis.Del(ctx, redisCategoryKey(id))
	s.redis.Del(ctx, redisCategoryListKey)
	return s.repo.delete(id)
}

func (s *topicCategoryServiceImpl) list(limit, offset int) ([]TopicCategory, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("%s:%d:%d", redisCategoryListKey, limit, offset)

	// Try cache
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var categories []TopicCategory
		if err := json.Unmarshal([]byte(cached), &categories); err == nil {
			return categories, nil
		}
	}

	// Fallback to DB
	categories, err := s.repo.list(limit, offset)
	if err != nil {
		return nil, err
	}

	// Cache
	if data, err := json.Marshal(categories); err == nil {
		s.redis.Set(ctx, cacheKey, data, categoryCacheTTL)
	}

	return categories, nil
}
/**************************
 * LIKE SERVICE INTERFACE *
 **************************/

type topicLikeService interface {
	create(like *TopicLike) error
	get(topicID, userID uuid.UUID) (*TopicLike, error)
	delete(topicID, userID uuid.UUID) error
	count(topicID uuid.UUID, likeValue bool) (int64, error)
}

type topicLikeServiceImpl struct {
	repo topicLikeRepository
	redis *redis.Client
}

func newTopicLikeService(repo topicLikeRepository, redis *redis.Client) topicLikeService {
	return &topicLikeServiceImpl{repo: repo, redis: redis}
}

func redisTopicLikeKey(topicID, userID uuid.UUID) string {
	return fmt.Sprintf("topiclike:%s:%s", topicID.String(), userID.String())
}

/**********************************
 * LIKE REPOSITORY IMPLEMENTATION *
 **********************************/

func (s *topicLikeServiceImpl) count(topicID uuid.UUID, likeValue bool) (int64, error) {
	return s.repo.count(topicID, likeValue)
}

func (s *topicLikeServiceImpl) create(like *TopicLike) error {
	if err := s.repo.create(like); err != nil {
		return err
	}

	// Cache in Redis
	key := redisTopicLikeKey(like.TopicID, like.UserID)
	data, _ := json.Marshal(like)
	_ = s.redis.Set(context.Background(), key, data, 10*time.Minute).Err()

	return nil
}

func (s *topicLikeServiceImpl) get(topicID, userID uuid.UUID) (*TopicLike, error) {
	ctx := context.Background()
	key := redisTopicLikeKey(topicID, userID)

	// Try Redis
	if cached, err := s.redis.Get(ctx, key).Result(); err == nil {
		var like TopicLike
		if err := json.Unmarshal([]byte(cached), &like); err == nil {
			return &like, nil
		}
	}

	// Fallback to DB
	like, err := s.repo.get(topicID, userID)
	if err != nil {
		return nil, err
	}

	// Cache result
	if data, err := json.Marshal(like); err == nil {
		_ = s.redis.Set(ctx, key, data, 10*time.Minute).Err()
	}

	return like, nil
}

func (s *topicLikeServiceImpl) delete(topicID, userID uuid.UUID) error {
	if err := s.repo.delete(topicID, userID); err != nil {
		return err
	}

	// Invalidate Redis
	key := redisTopicLikeKey(topicID, userID)
	_ = s.redis.Del(context.Background(), key).Err()

	return nil
}

/**********************************
 * COMMENT LIKE SERVICE INTERFACE *
 **********************************/

type commentLikeService interface {
	create(like *TopicCommentLike) error
	update(like *TopicCommentLike) error
	get(commentID, userID uuid.UUID) (*TopicCommentLike, error)
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

func (s *commentLikeServiceImpl) create(like *TopicCommentLike) error {
	if err := s.repo.create(like); err != nil {
		return err
	}
	// Cache it
	key := redisCommentLikeKey(like.CommentID, like.UserID)
	data, _ := json.Marshal(like)
	s.redis.Set(context.Background(), key, data, commentLikeTTL)
	return nil
}

func (s *commentLikeServiceImpl) update(like *TopicCommentLike) error {
	if err := s.repo.update(like); err != nil {
		return err
	}
	// Update cache
	key := redisCommentLikeKey(like.CommentID, like.UserID)
	data, _ := json.Marshal(like)
	s.redis.Set(context.Background(), key, data, commentLikeTTL)
	return nil
}

func (s *commentLikeServiceImpl) get(commentID, userID uuid.UUID) (*TopicCommentLike, error) {
	ctx := context.Background()
	key := redisCommentLikeKey(commentID, userID)

	// Try cache
	cached, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		var like TopicCommentLike
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
