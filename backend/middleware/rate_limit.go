package middleware

import (
	"fmt"
	"time"

	"pokemon/cache"

	"github.com/gofiber/fiber/v2"
)

const (
	Limit     = 10               // max requests
	Window    = time.Minute      // window duration
)
var RedisTTL = int(Window.Seconds())

func RateLimiter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		key := fmt.Sprintf("rate_limit:%s", ip)

		// Increment counter
		count, err := cache.RDB.Incr(cache.Ctx, key).Result()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Redis error")
		}

		// Set expiration only on first request
		if count == 1 {
			cache.RDB.Expire(cache.Ctx, key, Window)
		}

		// Check limit
		if count > Limit {
			ttl, _ := cache.RDB.TTL(cache.Ctx, key).Result()
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":       "Too many requests",
				"retry_after": int(ttl.Seconds()),
			})
		}

		return c.Next()
	}
}
