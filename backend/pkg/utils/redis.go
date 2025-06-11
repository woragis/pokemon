package utils

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func CacheKey(prefix string, teamID uuid.UUID) string {
	return fmt.Sprintf("%s:%s", prefix, teamID.String())
}

func GetCachedCount(redisClient *redis.Client, key string) (int64, error) {
	val, err := redisClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return -1, nil // Not in cache
	}
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(val, 10, 64)
}

func SetCachedCount(redisClient *redis.Client, key string, count int64) error {
	return redisClient.Set(context.Background(), key, count, time.Hour).Err()
}

func InvalidateCache(redisClient *redis.Client, key string) {
	redisClient.Del(context.Background(), key)
}
