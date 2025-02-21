package transporthttp_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/IordanisCodeExamples/go-api-example/internal/transport"
	transporthttp "github.com/IordanisCodeExamples/go-api-example/internal/transport/http"
)

func TestGetMovieHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mocks := getMocks(t)
		movie := "TheGodfather"
		mocks.service.
			EXPECT().
			GetMovie(gomock.Any(), gomock.Any()).
			Return(&transport.Movie{
				Title:            movie,
				Year:             1972,
				Duration:         175,
				Director:         "Francis Ford Coppola",
				Cast:             []string{"Marlon Brando", "Al Pacino", "James Caan"},
				Genre:            []string{"Crime", "Drama"},
				Synopsis:         "The aging patriarch of an organized crime dynasty.",
				BoxOfficeRevenue: 245066411,
			}, nil)

		server, err := transporthttp.NewHttpServer(mocks.ctx, mocks.service)
		assert.NoError(t, err)
		movieName := url.PathEscape(movie)
		req := httptest.NewRequest("GET", "/movies/"+movieName, nil)
		w := httptest.NewRecorder()
		server.GetMovieHandler(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("movie not found", func(t *testing.T) {
		mocks := getMocks(t)

		mocks.service.
			EXPECT().
			GetMovie(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("movie not found"))

		server, err := transporthttp.NewHttpServer(mocks.ctx, mocks.service)
		assert.NoError(t, err)
		req := httptest.NewRequest("GET", "/movies/NonExistentMovie", nil)
		w := httptest.NewRecorder()
		server.GetMovieHandler(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
