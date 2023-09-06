package transportkafka

import (
	"context"
	"fmt"
)

// Service is the interface that wraps the service layer methods
type Service interface {
	// IngestMovie ingest a movie to the database
	IngestMovie(ctx context.Context, movie Movie) error
}

// Consumer represents the kafka consumer
type Consumer struct {
	// Service is the service layer
	Service Service
}

func NewConsumer(service Service) (*Consumer, error) {
	if service == nil {
		return nil, fmt.Errorf("missing parameter service")
	}
	return &Consumer{
		Service: service,
	}, nil
}
