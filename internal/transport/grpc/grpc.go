package transportgrpc

import (
	"context"

	pb "github.com/junkd0g/go-api-example-schema/go/api"
	merror "github.com/junkd0g/neji"

	"github.com/IordanisCodeExamples/go-api-example/internal/transport"
)

// Service is the interface that wraps the service layer methods
type Service interface {
	GetMovie(ctx context.Context, title string) (*transport.Movie, error)
}

// HttpServer represents the transport's http server
type GrpcServer struct {
	ctx context.Context
	pb.UnimplementedMovieServiceServer
	Service Service
}

// NewHttpGprc creates a new http server
func NewGprcServer(
	ctx context.Context,
	service Service,
) (*GrpcServer, error) {
	if service == nil {
		return nil, merror.ErrInvalidParameter("service")
	}
	return &GrpcServer{
		ctx:     ctx,
		Service: service,
	}, nil
}
