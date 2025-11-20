package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string

	RedisAddr     string
	RedisPassword string
	RedisDB       int

	AppPort string
	Env     string
}

func NewConfig() *Config {
	_ = loadDotEnv()
	c := &Config{
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresDB:       getEnv("POSTGRES_DB", "app_db"),

		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),

		AppPort: getEnv("APP_PORT", "8080"),
		Env:     getEnv("ENV", "development"),
	}
	log.Printf("config loaded: env=%s, pg=%s@%s/%s", c.Env, c.PostgresUser, c.PostgresHost, c.PostgresDB)
	return c
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

func loadDotEnv() error {
	if _, ok := os.LookupEnv("ENV"); ok {
		return nil
	}
	if _, err := os.Stat(".env"); err == nil {
		return godotenv.Load()
	}
	return nil
}
