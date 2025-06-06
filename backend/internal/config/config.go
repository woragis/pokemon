package config

import (
	"os"
)

type Config struct {
    Port        string
    DatabaseURL string
    RedisURL    string
    JWTSecret   string
}

func Load() *Config {
    return &Config{
        Port:        getEnv("PORT", "3000"),
        DatabaseURL: getEnv("DATABASE_URL", "postgres://user:pass@localhost/dbname?sslmode=disable"),
        RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379"),
        JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
