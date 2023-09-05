package service_test

import (
	"testing"

	"github.com/junkd0g/go-api-example/internal/service"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Run("Creates successfully a Service object", func(t *testing.T) {
		// TODO
	})
	t.Run("Returns error when store is nil", func(t *testing.T) {
		srv, err := service.New(nil)
		assert.NotNil(t, err)
		assert.Nil(t, srv)
		assert.Contains(t, err.Error(), "missing parameter store")
		assert.Contains(t, err.Error(), "store")
	})
}
