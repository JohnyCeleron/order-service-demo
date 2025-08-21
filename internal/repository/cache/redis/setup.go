package redis

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func SetupRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
	return rdb, nil
}
