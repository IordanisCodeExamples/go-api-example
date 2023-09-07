package transportkafka

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	merror "github.com/junkd0g/neji"
)

// Service is the interface that wraps the service layer methods
type Service interface {
	// IngestMovie ingest a movie to the database
	IngestMovie(ctx context.Context, movie Movie) error
}

// Consumer represents the kafka consumer
type Consumer struct {
	Consumer          *kafka.Consumer
	Service           Service
	TopicsAndHandlers map[string]func(*kafka.Message) error
}

// NewConsumer creates a new kafka consumer
func NewConsumer(
	config *kafka.ConfigMap,
	service Service,
) (*Consumer, error) {
	if service == nil {
		return nil, merror.ErrInvalidParameter("service")
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka consumer: %w", err)
	}

	return &Consumer{
		Consumer: consumer,
		Service:  service,
	}, nil
}

// RegisterTopicHandlers registers the topic handlers
func (c *Consumer) RegisterTopicHandlers(topicsAndHandlers map[string]func(*kafka.Message) error) {
	c.TopicsAndHandlers = topicsAndHandlers
}

// Consume consumes messages from kafka and runs the handler provided
func (c *Consumer) Consume(handler func(*kafka.Message) error) {
	for {
		msg, err := c.Consumer.ReadMessage(-1)
		if err == nil {
			handler(msg)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

// StartConsuming starts consuming messages from kafka
func (c *Consumer) StartConsuming() {
	go c.Consume(func(msg *kafka.Message) error {
		handler, ok := c.TopicsAndHandlers[*msg.TopicPartition.Topic]
		if ok {
			fmt.Println("MARIKA2")
			handler(msg)
		} else {
			return fmt.Errorf("no handler found for topic %s", *msg.TopicPartition.Topic)
		}
		return nil
	})
}
