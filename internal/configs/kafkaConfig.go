package configs

type KafkaConfig struct {
	BootstrapServers    string
	GroupId             string
	DLQBootstrapServers string
	DLQTopic            string
	ConsumerTopic       string
	ProducerTopic       string
}

func NewKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		BootstrapServers:    getEnv("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092"),
		GroupId:             getEnv("KAFKA_GROUP_ID", "12341325"),
		DLQBootstrapServers: getEnv("KAFKA_DLQ_BOOTSTRAP_SERVERS", "localhost:9092"),
		DLQTopic:            getEnv("KAFKA_DLQ_TOPIC", "order-topic.dlq"),
		ConsumerTopic:       getEnv("KAFKA_CONSUMER_TOPIC", "order_topic"),
		ProducerTopic:       getEnv("KAFKA_PRODUCER_TOPIC", "order_topic"),
	}
}
