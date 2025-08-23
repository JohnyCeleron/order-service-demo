package order

import (
	"context"

	domainOrder "order-service/internal/domain/order"
)

type Closer interface {
	Close() error
}

type OrderRepositoryDB interface {
	Get(id string) (domainOrder.Order, error)
	GetAll() ([]domainOrder.Order, error)
	Add(order domainOrder.Order) error
	Closer
}

type OrderRepositoryCache interface {
	Get(ctx context.Context, key string) (domainOrder.Order, error)
	Set(ctx context.Context, key string, value domainOrder.Order) error
	Contains(ctx context.Context, id string) (bool, error)
	Closer
}

type OrderService struct {
	RepoDB         OrderRepositoryDB
	RepoCache      OrderRepositoryCache
	messageChannel chan domainOrder.Order
}

func NewService(repoDB OrderRepositoryDB, repoCache OrderRepositoryCache, messageChannel chan domainOrder.Order) *OrderService {
	return &OrderService{
		RepoDB:         repoDB,
		RepoCache:      repoCache,
		messageChannel: messageChannel,
	}
}
