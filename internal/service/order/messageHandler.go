package order

import (
	"context"

	"order-service/internal/app/logger"
	domainOrder "order-service/internal/domain/order"
	"order-service/internal/lib/logger/sl"
)

func (o *OrderService) HandleMessage(ctx context.Context, order domainOrder.Order) error {
	if err := o.RepoDB.Add(order); err != nil {
		logger.Logger.Error("handle message: ", sl.Err(err))
		return err
	}
	return nil
}
