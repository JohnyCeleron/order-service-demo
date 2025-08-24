package kafka

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func SetupKafkaConsumer() (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":           os.Getenv("KAFKA_GROUP_ID"),
		"enable.auto.commit": false,
		"auto.offset.reset":  "earliest",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	err = consumer.Subscribe(os.Getenv("KAFKA_CONSUMER_TOPIC"), nil)
	if err != nil {
		consumer.Close()
		return nil, fmt.Errorf("failed to subscribe to topic: %w", err)
	}

	return consumer, nil
}
