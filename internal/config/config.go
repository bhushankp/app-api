package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv        string
	HTTPPort      string
	DBDSN         string
	DBMaxOpen     int
	DBMaxIdle     int
	DBMaxLifetime time.Duration
}

func Load() *Config {
	// Load .env if present
	_ = godotenv.Load()

	cfg := &Config{
		AppEnv:        getEnvOrFail("APP_ENV"),
		HTTPPort:      getEnvOrFail("HTTP_PORT"),
		DBDSN:         getEnvOrFail("DB_DSN"),
		DBMaxOpen:     getEnvAsIntOrDefault("DB_MAX_OPEN", 25),
		DBMaxIdle:     getEnvAsIntOrDefault("DB_MAX_IDLE", 10),
		DBMaxLifetime: getEnvAsDurationOrDefault("DB_MAX_LIFETIME", 300*time.Second),
	}

	return cfg
}

func getEnvOrFail(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing required env var: %s", key)
	}
	return val
}

func getEnvAsIntOrDefault(key string, def int) int {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	var out int
	_, err := fmt.Sscanf(val, "%d", &out)
	if err != nil {
		return def
	}
	return out
}

func getEnvAsDurationOrDefault(key string, def time.Duration) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	dur, err := time.ParseDuration(val)
	if err != nil {
		return def
	}
	return dur
}
