package transportgrpc_test

import (
	"errors"
	"testing"

	pb "github.com/junkd0g/go-api-example-schema/go/api"
	"github.com/stretchr/testify/assert"

	"github.com/IordanisCodeExamples/go-api-example/internal/transport"
	transportgrpc "github.com/IordanisCodeExamples/go-api-example/internal/transport/grpc"
)

func Test_GetMovie(t *testing.T) {
	mocks := getMocks(t)
	server, _ := transportgrpc.NewGprcServer(mocks.ctx, mocks.service)

	t.Run("Returns movie successfully", func(t *testing.T) {
		expectedMovie := &transport.Movie{
			Title:            "Test Movie",
			Year:             2021,
			Duration:         120,
			Director:         "John Doe",
			Cast:             []string{"Jane Doe"},
			Genre:            []string{"Action"},
			Synopsis:         "Test Synopsis",
			BoxOfficeRevenue: 100,
		}
		mocks.service.EXPECT().GetMovie(mocks.ctx, "Test Movie").Return(expectedMovie, nil)

		req := &pb.MovieRequest{Title: "Test Movie"}
		resp, err := server.GetMovie(mocks.ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, expectedMovie.Title, resp.Title)
	})

	t.Run("Handles service error", func(t *testing.T) {
		mocks.service.EXPECT().GetMovie(mocks.ctx, "Error Movie").Return(nil, errors.New("service error"))

		req := &pb.MovieRequest{Title: "Error Movie"}
		resp, err := server.GetMovie(mocks.ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "service error")
	})
}
