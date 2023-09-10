package transportgrpc

import (
	"context"

	pb "github.com/junkd0g/go-api-example-schema/go/api"

	"github.com/junkd0g/go-api-example/internal/transport"
)

// GetMovie is the grpc handler for getting a movie
func (s *GrpcServer) GetMovie(ctx context.Context, req *pb.MovieRequest) (*pb.MovieResponse, error) {
	movie, err := s.Service.GetMovie(ctx, req.Title)
	if err != nil {
		return nil, err
	}

	return convertToMovieResponse(movie), nil
}

func convertToMovieResponse(movie *transport.Movie) *pb.MovieResponse {
	return &pb.MovieResponse{
		Title:            movie.Title,
		Year:             int32(movie.Year),
		Duration:         int32(movie.Duration),
		Director:         movie.Director,
		Cast:             movie.Cast,
		Genre:            movie.Genre,
		Synopsis:         movie.Synopsis,
		BoxOfficeRevenue: movie.BoxOfficeRevenue,
	}
}
