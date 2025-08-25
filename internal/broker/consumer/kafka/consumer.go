package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"order-service/internal/app/logger"
	domainOrder "order-service/internal/domain/order"
	"order-service/internal/lib/logger/sl"
	"order-service/internal/repository/db"
)

const (
	readingMessageTimeout time.Duration = 5 * time.Second
	timeoutFlushMs        int           = 10 * 1000
)

type MessageHandler interface {
	HandleMessage(ctx context.Context, order domainOrder.Order) error
}

type Consumer struct {
	reader         *kafka.Consumer
	messageHandler MessageHandler

	dlqProducer *kafka.Producer
}

func New(messageHandler MessageHandler) (*Consumer, error) {
	reader, err := SetupKafkaConsumer()
	if err != nil {
		return &Consumer{}, err
	}
	dlqProd, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_DLQ_BOOTSTRAP_SERVERS"),
	})
	if err != nil {
		reader.Close()
		return &Consumer{}, err
	}
	return &Consumer{
		reader:         reader,
		messageHandler: messageHandler,
		dlqProducer:    dlqProd,
	}, nil
}

func (c *Consumer) Run(ctx context.Context) {
	logger.Logger.Info("consumer runs")

	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := c.reader.ReadMessage(readingMessageTimeout)
			if msg == nil {
				continue
			}
			if err != nil {
				if kafkaErr, ok := err.(kafka.Error); ok && kafkaErr.Code() == kafka.ErrTimedOut {
					continue
				}
				logger.Logger.Warn("consumer reading warning: ", sl.Err(err))
				continue
			}

			var order domainOrder.Order
			if err = json.Unmarshal(msg.Value, &order); err != nil {
				if dlqErr := c.sendToDLQ(ctx, msg, fmt.Errorf("unmarshal: %w", err)); dlqErr != nil {
					logger.Logger.Error("dlq error: ", sl.Err(dlqErr))
				}
				c.reader.CommitMessage(msg)
				continue
			}
			if valid, err := order.Validate(); !valid {
				logger.Logger.Error("error validation: ", sl.Err(err))
				if dlqErr := c.sendToDLQ(ctx, msg, fmt.Errorf("validate: %w", err)); dlqErr != nil {
					logger.Logger.Error("dlq error: ", sl.Err(dlqErr))
				}
				c.reader.CommitMessage(msg)
				continue
			}
			if err := c.messageHandler.HandleMessage(ctx, order); err != nil {
				if errors.Is(err, db.ErrExistsKey) {
					logger.Logger.Error("order exists", slog.String("uid", order.OrderUID))
				}
				logger.Logger.Error("handle message error: ", sl.Err(err))
			}
			if _, err := c.reader.CommitMessage(msg); err != nil {
				logger.Logger.Warn("commit warn: ", sl.Err(err))
			}
		}
	}
}

func (c *Consumer) Close() error {
	c.dlqProducer.Flush(timeoutFlushMs)
	c.dlqProducer.Close()
	return c.reader.Close()
}
