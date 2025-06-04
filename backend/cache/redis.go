package cache

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)


var RDB *redis.Client
var Ctx = context.Background()

func InitRedis()  {
	// Get redis address
	addr := os.Getenv("REDIS_ADDRESS")
	if addr == "" {
		addr = "localhost:6379"
	}

	// Get redis password
	password := os.Getenv("REDIS_PASSWORD")

	// Get redis db
	dbStr := os.Getenv("REDIS_DB")
	db := 0 // Default DB
	if dbStr != "" {
		if parsed, err := strconv.Atoi(dbStr); err == nil {
			db = parsed
		}
	}
	
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Ping redis DB
	pong, err := RDB.Ping(Ctx).Result()
	if err != nil {
		panic("Redis connection failed: " + err.Error())
	}
	fmt.Println("Redis ping:", pong)
}
