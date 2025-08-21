package order

import (
	"context"

	"order-service/internal/domain"
)

var errorNotFound error

func (o *OrderService) GetById(ctx context.Context, id string) (domain.Order, error) {
	orderInCache, _ := o.repoCache.Contains(ctx, id)
	if orderInCache {
		order, err := o.repoCache.Get(ctx, id)
		if err == nil {
			return order, nil
		}
	}
	order, err := o.repoDB.Get(id)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
