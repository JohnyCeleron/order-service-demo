package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"order-service/internal/configs"
)

func SetupKafkaConsumer() (*kafka.Consumer, error) {
	cfg := configs.NewKafkaConfig()
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  cfg.BootstrapServers,
		"group.id":           cfg.GroupId,
		"enable.auto.commit": false,
		"auto.offset.reset":  "earliest",
		//"debug":              "consumer,cgrp,topic,fetch",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	err = consumer.Subscribe(cfg.ConsumerTopic, nil)
	if err != nil {
		consumer.Close()
		return nil, fmt.Errorf("failed to subscribe to topic: %w", err)
	}

	return consumer, nil
}
