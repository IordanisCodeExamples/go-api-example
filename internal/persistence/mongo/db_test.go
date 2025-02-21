package mongostore_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"

	store "github.com/IordanisCodeExamples/go-api-example/internal/persistence/mongo"
)

func Test_New(t *testing.T) {
	t.Run("Creates successfully a DB object", func(t *testing.T) {
		ctx := context.Background()
		mongo, _ := mongo.Connect(ctx, nil)
		db, err := store.New(mongo)
		assert.Nil(t, err)
		assert.NotNil(t, db)
	})

	t.Run("Returns error when mongo client is nil", func(t *testing.T) {
		db, err := store.New(nil)
		assert.NotNil(t, err)
		assert.Nil(t, db)
		assert.Contains(t, err.Error(), "missing parameter client")
		assert.Contains(t, err.Error(), "client")
	})
}
