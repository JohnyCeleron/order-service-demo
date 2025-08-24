package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	domainOrder "order-service/internal/domain/order"
)

const (
	readingMessageTimeout time.Duration = 2 * time.Second
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
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := c.reader.ReadMessage(readingMessageTimeout)
			if err != nil {
				if kafkaErr, ok := err.(kafka.Error); ok && kafkaErr.Code() == kafka.ErrTimedOut {
					continue
				}
				log.Printf("error reading from broker: %v\n", err)
				continue
			}

			var order domainOrder.Order
			if err = json.Unmarshal(msg.Value, &order); err != nil {
				if dlqErr := c.sendToDLQ(ctx, msg, fmt.Errorf("unmarshal: %w", err)); dlqErr != nil {
					log.Printf("dlq error: %v\n", dlqErr)
				}
				c.reader.CommitMessage(msg)
				continue
			}
			if valid, err := order.Validate(); !valid {
				if dlqErr := c.sendToDLQ(ctx, msg, fmt.Errorf("validate: %w", err)); dlqErr != nil {
					log.Printf("dlq error: %v\n", dlqErr)
				}
				c.reader.CommitMessage(msg)
				continue
			}

			if err := c.messageHandler.HandleMessage(ctx, order); err != nil {
				log.Printf("handle message error: %v\n")
			}
			if _, err := c.reader.CommitMessage(msg); err != nil {
				log.Printf("commit warn: %v", err)
			}
		}
	}
}

func (c *Consumer) Close() error {
	c.dlqProducer.Flush(timeoutFlushMs)
	c.dlqProducer.Close()
	return c.reader.Close()
}
