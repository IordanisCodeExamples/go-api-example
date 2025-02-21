package transportkafka_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	internalctx "github.com/IordanisCodeExamples/go-api-example/internal/context"
	internallogger "github.com/IordanisCodeExamples/go-api-example/internal/logger"
	transportkafkamock "github.com/IordanisCodeExamples/go-api-example/internal/mocks/transport/kafka"
	transport "github.com/IordanisCodeExamples/go-api-example/internal/transport"
	transportkafka "github.com/IordanisCodeExamples/go-api-example/internal/transport/kafka"
)

type mocks struct {
	ctx        context.Context
	service    *transportkafkamock.MockService
	configData *kafka.ConfigMap
}

func getMocks(t *testing.T) *mocks {
	t.Helper()
	logger, err := internallogger.NewLogger()
	assert.NoError(t, err)

	configMap := kafka.ConfigMap{
		"bootstrap.servers": "somehost",
		"group.id":          "somegroup",
	}

	assert.NoError(t, err)

	ctrl := gomock.NewController(t)
	service := transportkafkamock.NewMockService(ctrl)
	return &mocks{
		ctx:        internalctx.AddLoggerToContex(context.Background(), logger),
		service:    service,
		configData: &configMap,
	}
}

func Test_NewConsumer(t *testing.T) {
	mocks := getMocks(t)
	t.Run("Creates successfully a Consumer object", func(t *testing.T) {
		server, err := transportkafka.NewConsumer(mocks.ctx, mocks.configData, mocks.service)
		assert.Nil(t, err)
		assert.NotNil(t, server)
	})
	t.Run("Returns error when service is nil", func(t *testing.T) {
		server, err := transportkafka.NewConsumer(mocks.ctx, mocks.configData, nil)
		assert.NotNil(t, err)
		assert.Nil(t, server)
		assert.Contains(t, err.Error(), "missing parameter service")
		assert.Contains(t, err.Error(), "service")
	})

	t.Run("Returns error when config is nil", func(t *testing.T) {
		server, err := transportkafka.NewConsumer(mocks.ctx, nil, mocks.service)
		assert.NotNil(t, err)
		assert.Nil(t, server)
		assert.Contains(t, err.Error(), "missing parameter config")
		assert.Contains(t, err.Error(), "config")
	})
}

func Test_RegisterTopicHandlers(t *testing.T) {
	mocks := getMocks(t)

	t.Run("Successfully registers topic handlers", func(t *testing.T) {
		// Create a new Consumer
		consumer, err := transportkafka.NewConsumer(
			mocks.ctx,
			mocks.configData,
			mocks.service,
		)
		assert.Nil(t, err)
		assert.NotNil(t, consumer)

		// Define mock topic handlers
		topicsAndHandlers := map[string]func(*kafka.Message) error{
			"topic1": func(msg *kafka.Message) error { return nil },
			"topic2": func(msg *kafka.Message) error { return nil },
		}

		// Register topic handlers
		consumer.RegisterTopicHandlers(topicsAndHandlers)

		assert.Equal(t, topicsAndHandlers, consumer.TopicsAndHandlers)
	})
}

func Test_HandleInsertMovie(t *testing.T) {

	t.Run("Successfully decodes and ingests movie", func(t *testing.T) {
		// Create a valid Kafka message
		mocks := getMocks(t)
		movie := transport.Movie{Title: "TestMovie", Year: 2023}
		movieBytes, _ := json.Marshal(movie)
		mockMessage := &kafka.Message{
			Value: movieBytes,
		}

		mocks.service.EXPECT().IngestMovie(mocks.ctx, movie).Return(nil)

		consumer, err := transportkafka.NewConsumer(
			mocks.ctx,
			mocks.configData,
			mocks.service,
		)
		assert.Nil(t, err)

		err = consumer.HandleInsertMovie(mockMessage)
		assert.Nil(t, err)
	})

	t.Run("Returns error when failed to decode message", func(t *testing.T) {
		mocks := getMocks(t)

		mockMessage := &kafka.Message{
			Value: []byte("invalid json"),
		}

		consumer, err := transportkafka.NewConsumer(
			mocks.ctx,
			mocks.configData,
			mocks.service,
		)
		assert.Nil(t, err)

		err = consumer.HandleInsertMovie(mockMessage)
		assert.NotNil(t, err)
	})

	t.Run("Returns error when IngestMovie fails", func(t *testing.T) {
		mocks := getMocks(t)

		movie := transport.Movie{Title: "TestMovie", Year: 2023}
		movieBytes, _ := json.Marshal(movie)
		mockMessage := &kafka.Message{
			Value: movieBytes,
		}

		mocks.service.EXPECT().IngestMovie(mocks.ctx, movie).Return(errors.New("some error"))

		consumer, err := transportkafka.NewConsumer(
			mocks.ctx,
			mocks.configData,
			mocks.service,
		)
		assert.Nil(t, err)

		err = consumer.HandleInsertMovie(mockMessage)
		assert.NotNil(t, err)
		assert.Equal(t, "some error", err.Error())
	})
}
