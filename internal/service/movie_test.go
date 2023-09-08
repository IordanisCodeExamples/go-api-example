package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	mongostore "github.com/junkd0g/go-api-example/internal/persistence/mongo"
	"github.com/junkd0g/go-api-example/internal/service"
	"github.com/junkd0g/go-api-example/internal/transport"
)

func Test_IngestMovie(t *testing.T) {
	t.Run("Ingests a movie successfully", func(t *testing.T) {
		// arrange
		mocks := getMocks(t)

		mocks.store.
			EXPECT().
			InsertMovie(
				mocks.ctx,
				mongostore.Movie{
					Title:            "The Godfather",
					Year:             1972,
					Duration:         175,
					Director:         "Francis Ford Coppola",
					Cast:             []string{"Marlon Brando", "Al Pacino", "James Caan"},
					Genre:            []string{"Crime", "Drama"},
					Synopsis:         "The aging patriarch of an organized crime dynasty.",
					BoxOfficeRevenue: 245066411,
				},
			).Return(nil, nil)

		srv, err := service.New(mocks.store)
		assert.NoError(t, err)
		assert.NotNil(t, srv)

		// act
		err = srv.IngestMovie(mocks.ctx, transport.Movie{
			Title:            "The Godfather",
			Year:             1972,
			Duration:         175,
			Director:         "Francis Ford Coppola",
			Cast:             []string{"Marlon Brando", "Al Pacino", "James Caan"},
			Genre:            []string{"Crime", "Drama"},
			Synopsis:         "The aging patriarch of an organized crime dynasty.",
			BoxOfficeRevenue: 245066411,
		})

		// assert
		assert.NoError(t, err)
	})

	t.Run("Returns error when store returns error", func(t *testing.T) {
		// arrange
		mocks := getMocks(t)
		mocks.store.
			EXPECT().
			InsertMovie(
				mocks.ctx,
				mongostore.Movie{
					Title:            "The Godfather",
					Year:             1972,
					Duration:         175,
					Director:         "Francis Ford Coppola",
					Cast:             []string{"Marlon Brando", "Al Pacino", "James Caan"},
					Genre:            []string{"Crime", "Drama"},
					Synopsis:         "The aging patriarch of an organized crime dynasty.",
					BoxOfficeRevenue: 245066411,
				},
			).Return(nil, errors.New("sone-error"))

		srv, err := service.New(mocks.store)
		assert.NoError(t, err)
		assert.NotNil(t, srv)

		// act
		err = srv.IngestMovie(mocks.ctx, transport.Movie{
			Title:            "The Godfather",
			Year:             1972,
			Duration:         175,
			Director:         "Francis Ford Coppola",
			Cast:             []string{"Marlon Brando", "Al Pacino", "James Caan"},
			Genre:            []string{"Crime", "Drama"},
			Synopsis:         "The aging patriarch of an organized crime dynasty.",
			BoxOfficeRevenue: 245066411,
		})

		// assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error_inserting_movie")
	})
}

func Test_GetMovie(t *testing.T) {
	t.Run("Fetches a movie successfully", func(t *testing.T) {
		// arrange
		mocks := getMocks(t)

		mocks.store.
			EXPECT().
			FindMovie(
				mocks.ctx,
				"The Godfather",
			).Return(&mongostore.Movie{
			Title:            "The Godfather",
			Year:             1972,
			Duration:         175,
			Director:         "Francis Ford Coppola",
			Cast:             []string{"Marlon Brando", "Al Pacino", "James Caan"},
			Genre:            []string{"Crime", "Drama"},
			Synopsis:         "The aging patriarch of an organized crime dynasty.",
			BoxOfficeRevenue: 245066411,
		}, nil)

		srv, err := service.New(mocks.store)
		assert.NoError(t, err)
		assert.NotNil(t, srv)

		// act
		movie, err := srv.GetMovie(mocks.ctx, "The Godfather")

		// assert
		assert.NoError(t, err)
		assert.Equal(t, "The Godfather", movie.Title)
	})

	t.Run("Returns error when store returns error", func(t *testing.T) {
		// arrange
		mocks := getMocks(t)

		mocks.store.
			EXPECT().
			FindMovie(
				mocks.ctx,
				"The Godfather",
			).Return(nil, errors.New("some-error"))

		srv, err := service.New(mocks.store)
		assert.NoError(t, err)
		assert.NotNil(t, srv)

		// act
		_, err = srv.GetMovie(mocks.ctx, "The Godfather")

		// assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error_fetching_movie")
	})
}

func Test_DeleteMovie(t *testing.T) {
	t.Run("Deletes a movie successfully", func(t *testing.T) {
		// arrange
		mocks := getMocks(t)

		objectID := primitive.NewObjectID()
		deleteResult := &mongo.DeleteResult{DeletedCount: 1}

		mocks.store.
			EXPECT().
			DeleteMovie(mocks.ctx, objectID).
			Return(deleteResult, nil)

		srv, err := service.New(mocks.store)
		assert.NoError(t, err)
		assert.NotNil(t, srv)

		// act
		err = srv.DeleteMovie(mocks.ctx, objectID.Hex())

		// assert
		assert.NoError(t, err)
	})

	t.Run("Returns error when store returns error", func(t *testing.T) {
		// arrange
		mocks := getMocks(t)

		objectID := primitive.NewObjectID()

		mocks.store.
			EXPECT().
			DeleteMovie(mocks.ctx, objectID).
			Return(nil, errors.New("some-error"))

		srv, err := service.New(mocks.store)
		assert.NoError(t, err)
		assert.NotNil(t, srv)

		// act
		err = srv.DeleteMovie(mocks.ctx, objectID.Hex())

		// assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error_deleting_movie")
	})

	t.Run("Returns error when provided ID is invalid", func(t *testing.T) {
		// arrange
		mocks := getMocks(t)
		invalidID := "invalidObjectID"

		srv, err := service.New(mocks.store)
		assert.NoError(t, err)
		assert.NotNil(t, srv)

		// act
		err = srv.DeleteMovie(mocks.ctx, invalidID)

		// assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed_to_convert_string_to_objectid")
	})
}
