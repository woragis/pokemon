package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
    // Server
    Port string
    Env  string

    // Database
    DatabaseURL         string
    DatabaseMaxOpenConn int
    DatabaseMaxIdleConn int
    DatabaseMaxIdleTime time.Duration

    // Redis
    RedisURL      string
    RedisPassword string
    RedisDB       int

    // Security
    JWTSecret           string
    JWTExpireHours      int
    JWTRefreshExpireHours int
    AESKey              string
    HashSalt            string
    BCryptCost          int

    // Rate Limiting
    RateLimitMax    int
    RateLimitWindow time.Duration

    // CORS
    CORSAllowedOrigins []string
    CORSAllowedMethods []string
    CORSAllowedHeaders []string

    // Logging
    LogLevel  string
    LogFormat string

    // File Upload
    MaxFileSize int64
    UploadPath  string

    // Email
    SMTPHost     string
    SMTPPort     int
    SMTPUsername string
    SMTPPassword string
    EmailFrom    string

    // External APIs
    StripeSecretKey    string
    StripeWebhookSecret string
}

func Load() *Config {
    // Load .env file in development
    if os.Getenv("ENV") != "production" {
        if err := godotenv.Load(); err != nil {
            log.Println("No .env file found")
        }
    }

    return &Config{
        // Server
        Port: getEnv("PORT", "3000"),
        Env:  getEnv("ENV", "development"),

        // Database
        DatabaseURL:         getEnvRequired("DATABASE_URL"),
        DatabaseMaxOpenConn: getEnvAsInt("DATABASE_MAX_OPEN_CONNS", 25),
        DatabaseMaxIdleConn: getEnvAsInt("DATABASE_MAX_IDLE_CONNS", 25),
        DatabaseMaxIdleTime: getEnvAsDuration("DATABASE_MAX_IDLE_TIME", "15m"),

        // Redis
        RedisURL:      getEnvRequired("REDIS_URL"),
        RedisPassword: getEnv("REDIS_PASSWORD", ""),
        RedisDB:       getEnvAsInt("REDIS_DB", 0),

        // Security
        JWTSecret:             getEnvRequired("JWT_SECRET"),
        JWTExpireHours:        getEnvAsInt("JWT_EXPIRE_HOURS", 24),
        JWTRefreshExpireHours: getEnvAsInt("JWT_REFRESH_EXPIRE_HOURS", 168),
        AESKey:                getEnvRequired("AES_KEY"),
        HashSalt:              getEnvRequired("HASH_SALT"),
        BCryptCost:            getEnvAsInt("BCRYPT_COST", 12),

        // Rate Limiting
        RateLimitMax:    getEnvAsInt("RATE_LIMIT_MAX", 100),
        RateLimitWindow: getEnvAsDuration("RATE_LIMIT_WINDOW", "60s"),

        // CORS
        CORSAllowedOrigins: getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"*"}),
        CORSAllowedMethods: getEnvAsSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE"}),
        CORSAllowedHeaders: getEnvAsSlice("CORS_ALLOWED_HEADERS", []string{"*"}),

        // Logging
        LogLevel:  getEnv("LOG_LEVEL", "info"),
        LogFormat: getEnv("LOG_FORMAT", "json"),

        // File Upload
        MaxFileSize: getEnvAsInt64("MAX_FILE_SIZE", 10*1024*1024), // 10MB
        UploadPath:  getEnv("UPLOAD_PATH", "./uploads"),

        // Email
        SMTPHost:     getEnv("SMTP_HOST", ""),
        SMTPPort:     getEnvAsInt("SMTP_PORT", 587),
        SMTPUsername: getEnv("SMTP_USERNAME", ""),
        SMTPPassword: getEnv("SMTP_PASSWORD", ""),
        EmailFrom:    getEnv("EMAIL_FROM", ""),

        // External APIs
        StripeSecretKey:     getEnv("PAYMENT_STRIPE_SECRET_KEY", ""),
        StripeWebhookSecret: getEnv("PAYMENT_STRIPE_WEBHOOK_SECRET", ""),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvRequired(key string) string {
    value := os.Getenv(key)
    if value == "" {
        log.Fatalf("Environment variable %s is required", key)
    }
    return value
}

func getEnvAsInt(key string, defaultValue int) int {
    valueStr := getEnv(key, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
        return value
    }
    return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
    valueStr := getEnv(key, "")
    if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
        return value
    }
    return defaultValue
}

func getEnvAsDuration(key string, defaultValue string) time.Duration {
    valueStr := getEnv(key, defaultValue)
    if value, err := time.ParseDuration(valueStr); err == nil {
        return value
    }
    duration, _ := time.ParseDuration(defaultValue)
    return duration
}

func getEnvAsSlice(key string, defaultValue []string) []string {
    valueStr := getEnv(key, "")
    if valueStr == "" {
        return defaultValue
    }
    return strings.Split(valueStr, ",")
}
