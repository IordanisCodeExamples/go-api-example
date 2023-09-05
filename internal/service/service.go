package service

import (
	"context"

	merror "github.com/junkd0g/neji"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	mongostore "github.com/junkd0g/go-api-example/internal/persistence/mongo"
)

// Store is the interface that wraps the basic methods of the persistence layer
type Store interface {
	InsertMovie(ctx context.Context, movie mongostore.Movie) (*mongo.InsertOneResult, error)
	FindMovie(ctx context.Context, title string) (*mongostore.Movie, error)
	DeleteMovie(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
}

// Service is the struct that represents the service layer
type Service struct {
	Store Store
}

// New creates a new instance of the service layer
func New(store Store) (*Service, error) {
	if store == nil {
		return nil, merror.ErrInvalidParameter("store")
	}
	return &Service{
		Store: store,
	}, nil
}
