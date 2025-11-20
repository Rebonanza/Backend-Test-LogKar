package redis

import (
	"context"
	"fmt"

	"github.com/local/be-test-logkar/internal/config"
	goredis "github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) (*goredis.Client, error) {
	opt := &goredis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	}
	client := goredis.NewClient(opt)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("redis ping: %w", err)
	}
	return client, nil
}
