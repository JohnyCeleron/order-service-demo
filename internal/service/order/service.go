package order

import (
	"context"

	"order-service/internal/domain"
)

type OrderRepositoryDB interface {
	Get(id string) (domain.Order, error)
	GetAll() ([]domain.Order, error)
	Add(order domain.Order) error
}

type OrderRepositoryCache interface {
	Get(ctx context.Context, key string) (domain.Order, error)
	Set(ctx context.Context, key string, value domain.Order) error
	Contains(ctx context.Context, id string) (bool, error)
}

type OrderService struct {
	repoDB         OrderRepositoryDB
	repoCache      OrderRepositoryCache
	messageChannel chan domain.Order
}

func NewService(repoDB OrderRepositoryDB, repoCache OrderRepositoryCache, messageChannel chan domain.Order) *OrderService {
	return &OrderService{
		repoDB:         repoDB,
		repoCache:      repoCache,
		messageChannel: messageChannel,
	}
}
