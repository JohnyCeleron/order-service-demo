package redis

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"

	domainOrder "order-service/internal/domain/order"
)

type Redis struct {
	rdb *redis.Client
}

func New() (*Redis, error) {
	rdb, err := SetupRedis()
	if err != nil {
		return &Redis{}, err
	}
	return &Redis{
		rdb: rdb,
	}, nil
}

func (client *Redis) Set(ctx context.Context, key string, value domainOrder.Order) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.rdb.Set(ctx, key, data, 0).Err()
}

func (client *Redis) Contains(ctx context.Context, key string) (bool, error) {
	//TODO:
	return false, nil
}

func (client *Redis) Get(ctx context.Context, key string) (domainOrder.Order, error) {
	data, err := client.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return domainOrder.Order{}, err
	}
	var order domainOrder.Order
	if err = json.Unmarshal(data, &order); err != nil {
		return domainOrder.Order{}, err
	}
	return order, nil
}

func (client *Redis) Close() error {
	return client.rdb.Close()
}
