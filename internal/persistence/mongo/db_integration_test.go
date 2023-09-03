//go:build integration
// +build integration

package mongostore_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	mongostore "github.com/junkd0g/go-api-example/internal/persistence/mongo"
)

const (
	mongoURI = "mongodb://username:password@localhost:27017"
)

func cleanupDB(t *testing.T, client *mongo.Client, dbName string) {
	ctx := context.Background()
	collections, err := client.Database(dbName).ListCollectionNames(ctx, bson.D{})
	require.NoError(t, err)
	assert.NotNil(t, collections)
	for _, collection := range collections {
		res, err := client.Database(dbName).Collection(collection).DeleteMany(ctx, bson.M{})
		require.NotNil(t, res)
		require.NoError(t, err)
	}
}

func buildMongoStore(t *testing.T) (*mongo.Client, *mongostore.DB, func()) {
	t.Helper()
	ctx := context.Background()

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	require.NoError(t, err)
	err = mongoClient.Connect(ctx)
	require.NoError(t, err)

	store, err := mongostore.New(mongoClient)
	require.NoError(t, err)

	cleanupDB(t, mongoClient, mongostore.DBName)

	return mongoClient, store, func() {
		cleanupDB(t, mongoClient, mongostore.DBName)
		_ = mongoClient.Disconnect(context.Background())
	}
}
