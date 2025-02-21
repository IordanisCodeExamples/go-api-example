package transportgrpc

import (
	"context"

	pb "github.com/junkd0g/go-api-example-schema/go/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	internalctx "github.com/IordanisCodeExamples/go-api-example/internal/context"
	"github.com/IordanisCodeExamples/go-api-example/internal/transport"
)

// GetMovie is the grpc handler for getting a movie
func (s *GrpcServer) GetMovie(ctx context.Context, req *pb.MovieRequest) (*pb.MovieResponse, error) {
	logger, _ := internalctx.GetLoggerFromContext(s.ctx)
	logger.Info("getMovie grpc operation started")
	movie, err := s.Service.GetMovie(ctx, req.Title)
	if err != nil {
		logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, "Failed to retrieve movie: %v", err)
	}

	resp := convertToMovieResponse(movie)
	logger.Info("getMovie grpc operation ended")
	return resp, nil
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
