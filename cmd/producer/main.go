package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"order-service/internal/configs"
)

var kafkaCfg *configs.KafkaConfig

func main() {
	kafkaCfg = configs.NewKafkaConfig()

	config := &kafka.ConfigMap{
		"bootstrap.servers": kafkaCfg.BootstrapServers,
	}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("Ошибка создания producer: %v", err)
	}
	defer producer.Close()
	fmt.Println("producer создан и подключен к Kafka")
	topic := kafkaCfg.ProducerTopic
	dirPath := filepath.Join("cmd", "producer", "testModels")
	if err := CreateTopic(topic); err != nil {
		log.Fatalf("ошибка при создании топика: %v", err)
	}
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Ошибка доставки: %v", ev.TopicPartition.Error)
				} else {
					fmt.Printf("Доставлено: %s → partition %d\n",
						string(ev.Key), ev.TopicPartition.Partition)
				}
			}
		}
	}()

	for _, entry := range entries {
		if !entry.IsDir() {
			filePath := filepath.Join(dirPath, entry.Name())
			if err := sendJSONFile(producer, topic, filePath); err != nil {
				log.Printf("ошибка обработки файла %s: %v", entry.Name(), err)
				continue
			}

			fmt.Printf("Успешно обработан: %s\n", entry.Name())
			time.Sleep(3 * time.Second)
		}
	}
	producer.Flush(15 * 1000)
}

func sendJSONFile(producer *kafka.Producer, topic, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: data,
			Key:   []byte(filepath.Base(filePath)),
			Headers: []kafka.Header{
				{Key: "source_file", Value: []byte(filePath)},
				{Key: "processed_at", Value: []byte(time.Now().Format(time.RFC3339))},
			},
		}, nil)
}

func CreateTopic(topic string) error {
	log.Println("topic creation")
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": kafkaCfg.BootstrapServers,
	})
	if err != nil {
		log.Println("kafka admin client creation error: ", err)
		return err
	}
	defer admin.Close()
	results, err := admin.CreateTopics(context.Background(),
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     3,
			ReplicationFactor: 1,
		}},
	)
	if err != nil {
		log.Println("creation topic error: ", err)

		return err
	}
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError {
			if result.Error.Code() == kafka.ErrTopicAlreadyExists {
				return nil
			}
			return result.Error
		}
	}
	return nil
}
