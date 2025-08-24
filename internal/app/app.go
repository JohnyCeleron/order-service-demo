package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"order-service/internal/app/logger"
	"order-service/internal/broker/consumer/kafka"
	"order-service/internal/lib/logger/sl"
	"order-service/internal/repository/cache/redis"
	"order-service/internal/repository/db/postgres"
	"order-service/internal/service/order"
)

type Application struct {
	serviceOrder *order.OrderService
	consumer     *kafka.Consumer
}

func New() (*Application, error) {
	logger.SetupLogger(os.Getenv("ENVIRONMENT"))

	repoOrderDB, err := postgres.New()
	if err != nil {
		return &Application{}, err
	}
	repoOrderCache, err := redis.New()
	if err != nil {
		return &Application{}, err
	}
	serviceOrder := order.NewService(repoOrderDB, repoOrderCache)
	consumer, err := kafka.New(serviceOrder)
	if err != nil {
		return &Application{}, err
	}

	return &Application{
		serviceOrder: serviceOrder,
		consumer:     consumer,
	}, nil
}

func (a *Application) Run() {
	logger.Logger.Info(
		"running order service",
		slog.String("env", os.Getenv("ENVIRONMENT")),
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	a.serviceOrder.PreLoad(ctx)

	go a.consumer.Run(ctx)

	<-ctx.Done()
	a.Close()
}

func (a *Application) Close() {
	logger.Logger.Info("Gracefull shutdown")

	logger.Logger.Info("Stop Kafka")
	if err := a.consumer.Close(); err != nil {
		logger.Logger.Error("Stop consumer error: ", sl.Err(err))
	}
	logger.Logger.Info("Stop DataBase")
	if err := a.serviceOrder.RepoDB.Close(); err != nil {
		logger.Logger.Error("Stop database error: ", sl.Err(err))
	}
	logger.Logger.Info("Stop Cache")
	if err := a.serviceOrder.RepoCache.Close(); err != nil {
		logger.Logger.Error("Stop cache error: ", sl.Err(err))
	}
}
