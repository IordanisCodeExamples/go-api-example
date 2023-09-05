package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	mongostore "github.com/junkd0g/go-api-example/internal/persistence/mongo"
	"github.com/junkd0g/go-api-example/internal/service"
	transportkafka "github.com/junkd0g/go-api-example/internal/transport/kafka"
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
		err = srv.IngestMovie(mocks.ctx, transportkafka.Movie{
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
		err = srv.IngestMovie(mocks.ctx, transportkafka.Movie{
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
