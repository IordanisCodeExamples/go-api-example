package service

import (
	"context"
	"fmt"

	transportkafka "github.com/junkd0g/go-api-example/internal/transport/kafka"
)

// IngestMovie ingests a movie in the database through the service layer
func (s *Service) IngestMovie(ctx context.Context, movie transportkafka.Movie) error {
	_, err := s.Store.InsertMovie(ctx, fromKafkaOjectToMongoObject(movie))
	if err != nil {
		return fmt.Errorf("error_inserting_movie %v", err)
	}
	return nil
}
