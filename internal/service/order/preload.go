package order

import (
	"context"

	"order-service/internal/app/logger"
	"order-service/internal/lib/logger/sl"
)

func (o *OrderService) PreLoad(ctx context.Context) {
	logger.Logger.Info("start preload")
	orders, err := o.RepoDB.GetAll()
	if err != nil {
		logger.Logger.Error("preload: error getting all records from database: ", sl.Err(err))
		return
	}
	for _, order := range orders {
		if err = o.RepoCache.Set(ctx, order.OrderUID, order); err != nil {
			logger.Logger.Error("preload: error writing order %v in cache: ", order.OrderUID, sl.Err(err))
			continue
		}
	}
	logger.Logger.Info("end preload")
}
