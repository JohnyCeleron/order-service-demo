package kafka

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	domainOrder "order-service/internal/domain/order"
)

const readingMessageTimeout time.Duration = time.Second

type Consumer struct {
	reader         *kafka.Consumer
	messageChannel chan domainOrder.Order
}

func New(messageChannel chan domainOrder.Order) (*Consumer, error) {
	reader, err := SetupKafkaConsumer()
	if err != nil {
		return &Consumer{}, err
	}
	return &Consumer{
		reader:         reader,
		messageChannel: messageChannel,
	}, nil
}

func (c *Consumer) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		log.Println("stop consumer: ", ctx.Err())
		close(c.messageChannel)
		log.Println("stop connection with broker")
		c.reader.Close()
		wg.Add(1)
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := c.reader.ReadMessage(readingMessageTimeout)
			if err != nil {
				log.Println("error reading from broker: ", err)
				continue
			}
			var order domainOrder.Order
			if err = json.Unmarshal(msg.Value, &order); err != nil {
				log.Println("error unmarshal broker message: ", err)
				continue
			}
			if valid, err := order.Validate(); !valid {
				log.Println("некорректные данные из брокера: ", err)
				continue
			}
			c.messageChannel <- order
		}
	}
}
