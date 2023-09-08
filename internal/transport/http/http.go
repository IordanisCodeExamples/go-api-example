package transporthttp

import (
	"context"

	"github.com/junkd0g/go-api-example/internal/transport"
)

type Service interface {
	GetMovie(ctx context.Context, title string) (*transport.Movie, error)
}

type HttpServer struct {
}
