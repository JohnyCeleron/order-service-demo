package redis

import (
	"github.com/redis/go-redis/v9"

	"order-service/internal/configs"
)

func SetupRedis() (*redis.Client, error) {
	cfg := configs.NewRedisConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})
	return rdb, nil
}
