// Package mongostore is the mongoDB version of the persistence layer of the service
package mongostore

import (
	merror "github.com/junkd0g/neji"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	//DBName is the name of the database
	DBName = "lamovie"
	//collections names of the database
	movieCollection = "movie"
)

// DB represents the instance of the mongo store
type DB struct {
	client          *mongo.Client
	movieCollection *mongo.Collection
}

// New creates new object for DB struct
func New(client *mongo.Client) (*DB, error) {
	if client == nil {
		return nil, merror.ErrInvalidParameter("client")
	}
	return &DB{
		client:          client,
		movieCollection: client.Database(DBName).Collection(movieCollection),
	}, nil
}
