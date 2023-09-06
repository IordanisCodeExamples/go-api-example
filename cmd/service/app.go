package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"

	transportkafka "github.com/junkd0g/go-api-example/internal/transport/kafka"
)

func main() {
	setUpKafkaConsumer()
}

func setUpKafkaConsumer() {
	configMap := kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	}

	consumer, err := transportkafka.NewConsumer(&configMap, nil)
	if err != nil {
		panic(err)
	}

	topicHandlers := map[string]func(*kafka.Message) error{
		"topic-insert-movie": consumer.HandleInsertMovie,
	}

	consumer.RegisterTopicHandlers(topicHandlers)

	consumer.StartConsuming()

	select {}
}
