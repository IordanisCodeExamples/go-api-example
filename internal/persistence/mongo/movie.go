package mongostore

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/junkd0g/go-api-example/internal/persistence"
)

// InsertMovie inserts a movie in the database
func (db *DB) InsertMovie(ctx context.Context, movie Movie) (*mongo.InsertOneResult, error) {
	return db.movieCollection.InsertOne(ctx, movie)
}

// FindMovie finds a movie in the database
func (db *DB) FindMovie(ctx context.Context, title string) (*Movie, error) {
	var movie Movie
	err := db.movieCollection.FindOne(ctx, bson.M{"title": title}).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, persistence.ErrNotFound
		}
		return nil, fmt.Errorf("error fetching movie: %v", err)
	}
	return &movie, nil
}

// UpdateMovie updates a movie in the database
func (db *DB) UpdateMovie(ctx context.Context, id primitive.ObjectID, updatedMovie Movie) (*mongo.UpdateResult, error) {
	return db.movieCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{Key: "$set", Value: updatedMovie},
		},
	)
}

// DeleteMovie deletes a movie from the database
func (db *DB) DeleteMovie(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return db.movieCollection.DeleteOne(ctx, bson.M{"_id": id})
}
