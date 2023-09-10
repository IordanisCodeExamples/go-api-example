package transporthttp

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	merror "github.com/junkd0g/neji"

	internalctx "github.com/junkd0g/go-api-example/internal/context"
)

// Service is the interface that wraps the service layer methods
func (s *HttpServer) GetMovieHandler(writer http.ResponseWriter, request *http.Request) {
	logger, _ := internalctx.GetLoggerFromContext(s.ctx)

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")

	jsonBody, status := s.getMovie(request)
	writer.WriteHeader(status)
	_, err := writer.Write(jsonBody)
	if err != nil {
		logger.Error(err.Error())
	}
}

func (s *HttpServer) getMovie(request *http.Request) ([]byte, int) {
	logger, _ := internalctx.GetLoggerFromContext(s.ctx)
	logger.Info("getMovie operation started")

	title := mux.Vars(request)["title"]
	movie, err := s.Service.GetMovie(s.ctx, title)
	if err != nil {
		logger.Error(err.Error())
		resp, err := merror.SimpeErrorResponseWithStatus(http.StatusInternalServerError, err)
		if err != nil {
			logger.Error(err.Error())
		}
		return []byte(resp), http.StatusInternalServerError
	}

	jsonBody, err := json.Marshal(movie)
	if err != nil {
		logger.Error(err.Error())
		resp, err := merror.SimpeErrorResponseWithStatus(http.StatusInternalServerError, err)
		if err != nil {
			logger.Error(err.Error())
		}
		return []byte(resp), http.StatusInternalServerError
	}

	logger.Info("getMovie operation ended")
	return jsonBody, http.StatusOK
}
