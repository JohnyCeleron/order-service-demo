package order

import (
	"context"
	"log"
)

func (o *OrderService) PreLoad(ctx context.Context) {
	log.Println("start preload")
	orders, err := o.repoDB.GetAll()
	if err != nil {
		log.Println("preload: ошибка получения объектов из бд: ", err)
		return
	}
	for _, order := range orders {
		if err = o.repoCache.Set(ctx, order.OrderUID, order); err != nil {
			log.Println("preload: ошибка записи объекта %v в кэш: ", order.OrderUID, err)
			continue
		}
	}
	log.Println("end preload")
}
