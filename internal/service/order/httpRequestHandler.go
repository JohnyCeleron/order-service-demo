package order

import (
	"context"
	"errors"

	"order-service/internal/app/logger"
	domainOrder "order-service/internal/domain/order"
	"order-service/internal/lib/logger/sl"
	"order-service/internal/repository/db"
)

func (o *OrderService) GetById(ctx context.Context, id string) (domainOrder.Order, error) {
	orderInCache, _ := o.RepoCache.Contains(ctx, id)
	if orderInCache {
		order, err := o.RepoCache.Get(ctx, id)
		if err == nil {
			return order, nil
		}
	}
	order, err := o.RepoDB.Get(id)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return domainOrder.Order{}, ErrOrderNotFound
		}
		return domainOrder.Order{}, err
	}
	if err = o.RepoCache.Set(ctx, order.OrderUID, order); err != nil {
		logger.Logger.Error("error order adding in cache: ", sl.Err(err))
	}
	return order, nil
}
