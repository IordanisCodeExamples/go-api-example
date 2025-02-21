package transporthttp_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	internalctx "github.com/IordanisCodeExamples/go-api-example/internal/context"
	internallogger "github.com/IordanisCodeExamples/go-api-example/internal/logger"
	transporthttpmock "github.com/IordanisCodeExamples/go-api-example/internal/mocks/transport/http"
	transporthttp "github.com/IordanisCodeExamples/go-api-example/internal/transport/http"
)

type mocks struct {
	ctx     context.Context
	service *transporthttpmock.MockService
}

func getMocks(t *testing.T) *mocks {
	t.Helper()
	logger, err := internallogger.NewLogger()
	assert.NoError(t, err)

	ctrl := gomock.NewController(t)
	service := transporthttpmock.NewMockService(ctrl)
	return &mocks{
		ctx:     internalctx.AddLoggerToContex(context.Background(), logger),
		service: service,
	}
}

func Test_NewHttpServer(t *testing.T) {
	mocks := getMocks(t)
	t.Run("Creates successfully a HttpServer object", func(t *testing.T) {
		server, err := transporthttp.NewHttpServer(mocks.ctx, mocks.service)
		assert.Nil(t, err)
		assert.NotNil(t, server)
	})
	t.Run("Returns error when service is nil", func(t *testing.T) {
		server, err := transporthttp.NewHttpServer(mocks.ctx, nil)
		assert.NotNil(t, err)
		assert.Nil(t, server)
		assert.Contains(t, err.Error(), "missing parameter service")
		assert.Contains(t, err.Error(), "service")
	})
}
