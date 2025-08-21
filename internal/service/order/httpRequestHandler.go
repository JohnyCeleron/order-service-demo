package order

import (
	"context"

	domainOrder "order-service/internal/domain/order"
)

var errorNotFound error

func (o *OrderService) GetById(ctx context.Context, id string) (domainOrder.Order, error) {
	orderInCache, _ := o.repoCache.Contains(ctx, id)
	if orderInCache {
		order, err := o.repoCache.Get(ctx, id)
		if err == nil {
			return order, nil
		}
	}
	order, err := o.repoDB.Get(id)
	if err != nil {
		return domainOrder.Order{}, err
	}
	return order, nil
}
