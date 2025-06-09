package forum

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ForumTopicRepository interface {
	Create(ctx context.Context, topic *Topic) error
	GetByID(ctx context.Context, id string) (*Topic, error)
	Update(ctx context.Context, id string, updates map[string]interface{}) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*Topic, error)
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewTopicRepository(db *gorm.DB, redis *redis.Client) ForumTopicRepository {
	return &repository{
		db:    db,
		redis: redis,
	}
}

func (r *repository) Create(ctx context.Context, topic *Topic) error {
	return r.db.WithContext(ctx).Create(topic).Error
}

func (r *repository) GetByID(ctx context.Context, id string) (*Topic, error) {
	cacheKey := r.getCacheKey(id)
	cached, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var topic Topic
		if json.Unmarshal([]byte(cached), &topic) == nil {
			return &topic, nil
		}
	}

	var topic Topic
	err = r.db.WithContext(ctx).Preload("Author").First(&topic, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	r.cacheTopic(ctx, &topic)
	return &topic, nil
}

func (r *repository) Update(ctx context.Context, id string, updates map[string]interface{}) error {
	err := r.db.WithContext(ctx).Model(&Topic{}).Where("id = ?", id).Updates(updates).Error
	if err == nil {
		r.redis.Del(ctx, r.getCacheKey(id))
	}
	return err
}

func (r *repository) Delete(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Delete(&Topic{}, "id = ?", id).Error
	if err == nil {
		r.redis.Del(ctx, r.getCacheKey(id))
	}
	return err
}

func (r *repository) List(ctx context.Context, limit, offset int) ([]*Topic, error) {
	var topics []*Topic
	err := r.db.WithContext(ctx).
		Preload("Author").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&topics).Error
	return topics, err
}

func (r *repository) getCacheKey(id string) string {
	return fmt.Sprintf("topic:%s", id)
}

func (r *repository) cacheTopic(ctx context.Context, topic *Topic) {
	data, _ := json.Marshal(topic)
	r.redis.Set(ctx, r.getCacheKey(topic.ID.String()), data, 15*time.Minute)
}
