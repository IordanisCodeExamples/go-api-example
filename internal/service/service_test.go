package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	servicemock "github.com/junkd0g/go-api-example/internal/mocks/service"
	"github.com/junkd0g/go-api-example/internal/service"
)

type serviceMocks struct {
	ctx   context.Context
	store *servicemock.MockStore
}

func getMocks(t *testing.T) serviceMocks {
	t.Helper()
	ctrl := gomock.NewController(t)

	return serviceMocks{
		ctx:   context.Background(),
		store: servicemock.NewMockStore(ctrl),
	}
}

func Test_New(t *testing.T) {
	mocks := getMocks(t)
	t.Run("Creates successfully a Service object", func(t *testing.T) {
		srv, err := service.New(mocks.store)
		assert.Nil(t, err)
		assert.NotNil(t, srv)
	})
	t.Run("Returns error when store is nil", func(t *testing.T) {
		srv, err := service.New(nil)
		assert.NotNil(t, err)
		assert.Nil(t, srv)
		assert.Contains(t, err.Error(), "missing parameter store")
		assert.Contains(t, err.Error(), "store")
	})
}
