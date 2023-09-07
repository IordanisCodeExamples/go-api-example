package service

import (
	"context"
	"fmt"

	transportkafka "github.com/junkd0g/go-api-example/internal/transport/kafka"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IngestMovie ingests a movie in the database through the service layer
func (s *Service) IngestMovie(ctx context.Context, movie transportkafka.Movie) error {
	_, err := s.Store.InsertMovie(ctx, fromKafkaOjectToMongoObject(movie))
	if err != nil {
		return fmt.Errorf("error_inserting_movie %v", err)
	}
	return nil
}

// GetMovie gets a movie from the database through the service layer
func (s *Service) GetMovie(
	ctx context.Context,
	title string,
) (*transportkafka.Movie, error) {
	movie, err := s.Store.FindMovie(ctx, title)
	if err != nil {
		return nil, fmt.Errorf("error_fetching_movie %v", err)
	}
	return fromMongoObjectToKafkaObject(*movie), nil
}

// DeleteMovie deletes a movie from the database through the service layer
func (s *Service) DeleteMovie(
	ctx context.Context,
	id string,
) error {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed_to_convert_string_to_objectid: %v", err)
	}

	_, err = s.Store.DeleteMovie(ctx, primitiveID)
	if err != nil {
		return fmt.Errorf("error_deleting_movie %v", err)
	}
	return nil
}
