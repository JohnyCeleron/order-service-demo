package order

import (
	"context"
	"log"

	domainOrder "order-service/internal/domain/order"
)

func (o *OrderService) HandleMessage(ctx context.Context, order domainOrder.Order) error {
	if err := o.RepoDB.Add(order); err != nil {
		log.Println("handle message: ", err)
		return err
	}
	return nil
}
