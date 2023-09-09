package transporthttp

import (
	"context"

	"github.com/gorilla/mux"
	merror "github.com/junkd0g/neji"

	"github.com/junkd0g/go-api-example/internal/transport"
)

// Service is the interface that wraps the service layer methods
type Service interface {
	GetMovie(ctx context.Context, title string) (*transport.Movie, error)
}

// HttpServer represents the transport's http server
type HttpServer struct {
	ctx     context.Context
	Service Service
}

// NewHttpServer creates a new http server
func NewHttpServer(
	ctx context.Context,
	service Service,
) (*HttpServer, error) {
	if service == nil {
		return nil, merror.ErrInvalidParameter("service")
	}

	return &HttpServer{
		ctx:     ctx,
		Service: service,
	}, nil
}

// GetRouter returns the router of the http server transport layer
// with the handlers registered
func (s *HttpServer) GetRouter() *mux.Router {
	var router = mux.NewRouter()
	router.HandleFunc("/movies/{title}", s.GetMovieHandler).Methods("GET")
	return router
}
