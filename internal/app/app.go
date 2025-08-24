package app

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"order-service/internal/broker/consumer/kafka"
	"order-service/internal/repository/cache/redis"
	"order-service/internal/repository/db/postgres"
	"order-service/internal/service/order"
)

type Application struct {
	serviceOrder *order.OrderService
	consumer     *kafka.Consumer
}

const messageBuffer int = 4

func New() (*Application, error) {
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

func (a *Application) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	a.serviceOrder.PreLoad(ctx)

	go a.consumer.Run(ctx)

	<-ctx.Done()
	a.Close()
	return nil
}

func (a *Application) Close() {
	log.Println("Gracefull Shutdown")

	log.Println("Stop Kafka")
	if err := a.consumer.Close(); err != nil {
		log.Println("Stop consumer error: ", err)
	}
	log.Println("Stop DataBase")
	if err := a.serviceOrder.RepoDB.Close(); err != nil {
		log.Println("Stop database error: ", err)
	}
	log.Println("Stop Cache")
	if err := a.serviceOrder.RepoCache.Close(); err != nil {
		log.Println("Stop cache error: ", err)
	}
}
