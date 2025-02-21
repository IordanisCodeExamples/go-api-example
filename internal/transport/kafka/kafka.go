package transportkafka

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	merror "github.com/junkd0g/neji"

	"github.com/IordanisCodeExamples/go-api-example/internal/transport"
)

// Service is the interface that wraps the service layer methods
type Service interface {
	// IngestMovie ingest a movie to the database
	IngestMovie(ctx context.Context, movie transport.Movie) error
}

// Consumer represents the kafka consumer
type Consumer struct {
	ctx               context.Context
	Consumer          *kafka.Consumer
	Service           Service
	TopicsAndHandlers map[string]func(*kafka.Message) error
}

// NewConsumer creates a new kafka consumer
func NewConsumer(
	ctx context.Context,
	config *kafka.ConfigMap,
	service Service,
) (*Consumer, error) {
	if service == nil {
		return nil, merror.ErrInvalidParameter("service")
	}

	if config == nil {
		return nil, merror.ErrInvalidParameter("config")
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka consumer: %w", err)
	}

	return &Consumer{
		ctx:      ctx,
		Consumer: consumer,
		Service:  service,
	}, nil
}

// RegisterTopicHandlers registers the topic handlers
func (c *Consumer) RegisterTopicHandlers(topicsAndHandlers map[string]func(*kafka.Message) error) {
	c.TopicsAndHandlers = topicsAndHandlers

	// Extract topic names from the map keys
	var topics []string
	for topic := range topicsAndHandlers {
		topics = append(topics, topic)
	}

	// Subscribe to the extracted topics
	c.Consumer.SubscribeTopics(topics, nil)
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
func (c *Consumer) StartConsuming(context.Context) {
	go c.Consume(func(msg *kafka.Message) error {
		handler, ok := c.TopicsAndHandlers[*msg.TopicPartition.Topic]
		if ok {
			handler(msg)
		} else {
			return fmt.Errorf("no handler found for topic %s", *msg.TopicPartition.Topic)
		}
		return nil
	})
}
