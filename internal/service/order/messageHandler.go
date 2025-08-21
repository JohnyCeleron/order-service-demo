package order

import (
	"context"
	"log"
	"sync"
)

func (o *OrderService) HandleMessage(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		wg.Add(1)
	}()

	for order := range o.messageChannel {
		if err := o.repoDB.Add(order); err != nil {
			log.Println("ошибка добавления заказа в бд: ", err)
		}
	}
}
