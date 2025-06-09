package user

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repository interface {
    Create(ctx context.Context, user *User) error
    GetByID(ctx context.Context, id uint) (*User, error)
    GetByEmail(ctx context.Context, email string) (*User, error)
    Update(ctx context.Context, id uint, updates map[string]interface{}) error
    Delete(ctx context.Context, id uint) error
    List(ctx context.Context, limit, offset int) ([]*User, error)
}

type repository struct {
    db    *gorm.DB
    redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) Repository {
    return &repository{
        db:    db,
        redis: redis,
    }
}

func (r *repository) Create(ctx context.Context, user *User) error {
    return r.db.WithContext(ctx).Create(user).Error
}

func (r *repository) GetByID(ctx context.Context, id uint) (*User, error) {
    // Try to get from cache first
    cacheKey := r.getUserCacheKey(id)
    cached, err := r.redis.Get(ctx, cacheKey).Result()
    if err == nil {
        var user User
        if json.Unmarshal([]byte(cached), &user) == nil {
            return &user, nil
        }
    }
    
    var user User
    err = r.db.WithContext(ctx).First(&user, id).Error
    if err != nil {
        return nil, err
    }
    
    // Cache the result
    r.cacheUser(ctx, &user)
    return &user, nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*User, error) {
    var user User
    err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
    return &user, err
}

func (r *repository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
    err := r.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Updates(updates).Error
    if err == nil {
        // Invalidate cache
        r.redis.Del(ctx, r.getUserCacheKey(id))
    }
    return err
}

func (r *repository) Delete(ctx context.Context, id uint) error {
    err := r.db.WithContext(ctx).Delete(&User{}, id).Error
    if err == nil {
        // Invalidate cache
        r.redis.Del(ctx, r.getUserCacheKey(id))
    }
    return err
}

func (r *repository) List(ctx context.Context, limit, offset int) ([]*User, error) {
    var users []*User
    err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&users).Error
    return users, err
}

func (r *repository) getUserCacheKey(id uint) string {
    return fmt.Sprintf("user:%d", id)
}

func (r *repository) cacheUser(ctx context.Context, user *User) {
    data, _ := json.Marshal(user)
    r.redis.Set(ctx, r.getUserCacheKey(user.ID), data, 15*time.Minute)
}
