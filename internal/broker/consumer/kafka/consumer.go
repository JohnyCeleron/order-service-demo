package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"order-service/internal/domain"
)

const readingMessageTimeout time.Duration = time.Second

type Consumer struct {
	reader         *kafka.Consumer
	messageChannel chan domain.Order
}

func New(messageChannel chan domain.Order) (*Consumer, error) {
	reader, err := SetupKafkaConsumer()
	if err != nil {
		return &Consumer{}, err
	}
	return &Consumer{
		reader:         reader,
		messageChannel: messageChannel,
	}, nil
}

func (c *Consumer) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("stop consumer: ", ctx.Err())
			close(c.messageChannel)
			c.reader.Close()
			return
		default:
			msg, err := c.reader.ReadMessage(readingMessageTimeout)
			if err != nil {
				log.Println("error reading from broker: ", err)
				continue
			}
			var order domain.Order
			if err = json.Unmarshal(msg.Value, &order); err != nil {
				log.Println("error unmarshal broker message: ", err)
				continue
			}
			c.messageChannel <- order
		}
	}
}
