package kafka

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const timeoutDLQ time.Duration = 5 * time.Second

func (c *Consumer) sendToDLQ(ctx context.Context, orig *kafka.Message, cause error) error {
	headers := getHeaders(cause, orig)
	dlqTopic := os.Getenv("KAFKA_DLQ_TOPIC")
	dlqMsg := getDLQMessage(dlqTopic, orig, headers)
	delivery := make(chan kafka.Event, 1)

	if err := c.dlqProducer.Produce(dlqMsg, delivery); err != nil {
		return fmt.Errorf("DLQ produce: %w", err)
	}
	select {
	case event := <-delivery:
		if msg, ok := event.(*kafka.Message); ok {
			if msg.TopicPartition.Error != nil {
				return fmt.Errorf("DLQ delivery failed: %w", msg.TopicPartition.Error)
			}
		}
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(timeoutDLQ):
		return fmt.Errorf("DLQ delivery timeout")
	}
	return nil
}

func getDLQMessage(dlqTopic string, orig *kafka.Message, headers []kafka.Header) *kafka.Message {
	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &dlqTopic,
			Partition: kafka.PartitionAny,
		},
		Key:     orig.Key,
		Value:   orig.Value,
		Headers: headers,
	}
}

func getHeaders(cause error, orig *kafka.Message) []kafka.Header {
	return []kafka.Header{
		{Key: "dlq-error", Value: []byte(cause.Error())},
		{Key: "orig-topic", Value: []byte(*orig.TopicPartition.Topic)},
		{Key: "orig-partition", Value: []byte(fmt.Sprint(orig.TopicPartition.Partition))},
		{Key: "orig-offset", Value: []byte(fmt.Sprint(orig.TopicPartition.Offset))},
	}
}
