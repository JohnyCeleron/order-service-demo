package app

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func newRedis() (*redis.Client, *context.Context, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return &redis.Client{}, nil, fmt.Errorf("ошибка подключения к Redis: %v", err)
	}
	return rdb, &ctx, nil
}
