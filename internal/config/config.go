package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv         string
	HTTPPort       string
	DB_DSN         string
	DB_MaxOpen     int
	DB_MaxIdle     int
	DB_MaxLifeTime time.Duration
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppEnv:         mustGet("APP_ENV"),
		HTTPPort:       mustGet("HTTP_PORT"),
		DB_DSN:         mustGet("DB_DSN"),
		DB_MaxOpen:     mustGetInt("DB_MAX_OPEN"),
		DB_MaxIdle:     mustGetInt("DB_MAX_IDLE"),
		DB_MaxLifeTime: mustGetDuration("DB_MAX_LIFETIME"),
	}
	return cfg
}

func mustGet(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("missing required env : %s", key)
	}
	return val
}

func mustGetInt(key string) int {
	val := mustGet(key)
	i, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("invalid int for %s: %v", key, err)
	}
	return i
}

func mustGetDuration(key string) time.Duration {
	val := mustGet(key)
	t, err := time.ParseDuration(val)
	if err != nil {
		log.Fatalf("invald duration for %s: %v", key, err)
	}
	return t
}
