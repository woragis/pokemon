package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedis(url string) *redis.Client {
    opt, _ := redis.ParseURL(url)
    RDB := redis.NewClient(opt)

	// Ping redis DB
	pong, err := RDB.Ping(Ctx).Result()
	if err != nil {
		panic("Redis connection failed: " + err.Error())
	}
	fmt.Println("Redis ping:", pong)

    return RDB
}
