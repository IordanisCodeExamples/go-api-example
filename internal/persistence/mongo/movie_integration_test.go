//go:build integration
// +build integration

package mongostore_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/junkd0g/go-api-example/internal/persistence"
	mongostore "github.com/junkd0g/go-api-example/internal/persistence/mongo"
)

func Test_InsertMovie(t *testing.T) {
	t.Run("insert movie successfully", func(t *testing.T) {
		ctx := context.Background()
		c, s, td := buildMongoStore(t)
		defer td()
		assert.NotNil(t, c)

		insertedMovie, err := s.InsertMovie(ctx, mongostore.Movie{
			Title:            "ExampleTitle2",
			Year:             2023,
			Duration:         120,
			Director:         "John Doe",
			Cast:             []string{"Actor 1", "Actor 2"},
			Genre:            []string{"Action", "Thriller"},
			Synopsis:         "An example movie synopsis",
			BoxOfficeRevenue: 32.5,
		})

		assert.Nil(t, err)
		assert.NotNil(t, insertedMovie)
	})
	t.Run("insert movie fails with canceled context", func(t *testing.T) {
		// Create a new context and immediately cancel it
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		c, s, td := buildMongoStore(t)
		defer td()
		assert.NotNil(t, c)

		insertedMovie, err := s.InsertMovie(ctx, mongostore.Movie{
			Title:            "Example Title",
			Year:             2023,
			Duration:         120,
			Director:         "John Doe",
			Cast:             []string{"Actor 1", "Actor 2"},
			Genre:            []string{"Action", "Thriller"},
			Synopsis:         "An example movie synopsis",
			BoxOfficeRevenue: 32.5,
		})

		// The insert should fail due to the canceled context
		assert.NotNil(t, err)
		assert.Nil(t, insertedMovie)
	})
}

func Test_FindMovie(t *testing.T) {
	t.Run("find movie successfully", func(t *testing.T) {
		ctx := context.Background()
		_, s, td := buildMongoStore(t)
		defer td()

		_, err := s.InsertMovie(ctx, mongostore.Movie{
			Title:            "Example Title",
			Year:             2023,
			Duration:         120,
			Director:         "John Doe",
			Cast:             []string{"Actor 1", "Actor 2"},
			Genre:            []string{"Action", "Thriller"},
			Synopsis:         "An example movie synopsis",
			BoxOfficeRevenue: 32.5,
		})

		assert.Nil(t, err)

		movieTitle := "Example Title"
		foundMovie, err := s.FindMovie(ctx, movieTitle)

		assert.Nil(t, err)
		assert.NotNil(t, foundMovie)
		assert.Equal(t, movieTitle, foundMovie.Title)
	})

	t.Run("find movie not found", func(t *testing.T) {
		ctx := context.Background()
		_, s, td := buildMongoStore(t)
		defer td()

		foundMovie, err := s.FindMovie(ctx, "Nonexistent Title")

		assert.Equal(t, persistence.ErrNotFound, err)
		assert.Nil(t, foundMovie)
	})

	t.Run("find movie fails with canceled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, s, td := buildMongoStore(t)
		defer td()

		foundMovie, err := s.FindMovie(ctx, "Example Title")

		assert.NotNil(t, err)
		assert.Nil(t, foundMovie)
	})
}

func Test_DeleteMovie(t *testing.T) {
	t.Run("delete movie successfully", func(t *testing.T) {
		ctx := context.Background()
		_, s, td := buildMongoStore(t)
		defer td()

		insertedMovie, err := s.InsertMovie(ctx, mongostore.Movie{
			Title:            "Example Title",
			Year:             2023,
			Duration:         120,
			Director:         "John Doe",
			Cast:             []string{"Actor 1", "Actor 2"},
			Genre:            []string{"Action", "Thriller"},
			Synopsis:         "An example movie synopsis",
			BoxOfficeRevenue: 32.5,
		})

		deleteResult, err := s.DeleteMovie(ctx, insertedMovie.InsertedID.(primitive.ObjectID))

		assert.Nil(t, err)
		assert.NotNil(t, deleteResult)
		assert.Equal(t, int64(1), deleteResult.DeletedCount)
	})

	t.Run("delete movie fails with canceled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, s, td := buildMongoStore(t)
		defer td()

		// Usage
		hexString, err := randomHex(12) // 12 bytes = 24 characters in hex
		if err != nil {
			t.Fatalf("Failed to generate hex string: %v", err)
		}

		movieID, err := primitive.ObjectIDFromHex(hexString)
		if err != nil {
			t.Fatalf("Failed to convert string to ObjectID: %v", err)
		}

		deleteResult, err := s.DeleteMovie(ctx, movieID)

		assert.NotNil(t, err)
		assert.Nil(t, deleteResult)
	})
}
