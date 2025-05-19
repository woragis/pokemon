package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables.")
    }
}

func GetDatabaseURL() string {
    url := os.Getenv("DATABASE_URL")
    if url == "" {
        log.Fatal("DATABASE_URL not set in environment")
    }
    return url
}
