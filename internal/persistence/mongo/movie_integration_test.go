//go:build integration
// +build integration

package mongostore_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	mongostore "github.com/junkd0g/go-api-example/internal/persistence/mongo"
)

func Test_InsertMovie(t *testing.T) {
	t.Run("insert movie successfully", func(t *testing.T) {
		ctx := context.Background()
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
