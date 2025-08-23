package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"order-service/internal/broker/consumer/kafka"
	domainOrder "order-service/internal/domain/order"
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
	messageChan := make(chan domainOrder.Order, messageBuffer)
	consumer, err := kafka.New(messageChan)
	if err != nil {
		return &Application{}, err
	}
	serviceOrder := order.NewService(repoOrderDB, repoOrderCache, messageChan)
	return &Application{
		serviceOrder: serviceOrder,
		consumer:     consumer,
	}, nil
}

func (a *Application) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		cancel()
	}()

	a.serviceOrder.PreLoad(ctx)

	var wg sync.WaitGroup

	wg.Add(2)

	go a.consumer.Run(ctx, &wg)
	go a.serviceOrder.HandleMessage(ctx, &wg)

	wg.Wait()

	<-ctx.Done()

	log.Println("Gracefull Shutdown")
	return nil
}
