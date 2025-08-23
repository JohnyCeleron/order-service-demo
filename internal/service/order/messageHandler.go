package order

import (
	"context"
	"log"
	"sync"
)

func (o *OrderService) HandleMessage(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range o.messageChannel {
		if err := o.RepoDB.Add(order); err != nil {
			log.Println("ошибка добавления заказа в бд: ", err)
		}
	}
}
