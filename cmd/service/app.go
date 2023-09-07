package main

import (
	"context"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/junkd0g/go-api-example/internal/config"
	internalctx "github.com/junkd0g/go-api-example/internal/context"
	internallogger "github.com/junkd0g/go-api-example/internal/logger"
	store "github.com/junkd0g/go-api-example/internal/persistence/mongo"
	"github.com/junkd0g/go-api-example/internal/service"
	transportkafka "github.com/junkd0g/go-api-example/internal/transport/kafka"
)

var (
	env        = os.Getenv("env")
	configPath = os.Getenv("configPath")
)

func main() {
	if len(env) == 0 {
		env = "dev"
		configPath = "./assets/config/dev.yaml"
	}
	ctx := context.Background()

	logger, err := internallogger.NewLogger()
	if err != nil {
		logger.Error("Error creating logger", internallogger.LogField{"error": err.Error()})
		panic(err)
	}

	ctx = internalctx.AddLoggerToContex(ctx, logger)

	configData, err := config.GetAppConfig(configPath)
	if err != nil {
		logger.Error("Error creating config", internallogger.LogField{"error": err.Error()})
		panic(fmt.Errorf("creating_config %w", err))
	}

	logger.Info("Starting Persistence")
	persistence, err := setUpPersistence(ctx, configData, env)
	if err != nil {
		logger.Error("Error creating persistence", internallogger.LogField{"error": err.Error()})
		panic(fmt.Errorf("creating_persistence %w", err))
	}

	srv, err := service.New(persistence)
	if err != nil {
		logger.Error("Error creating service", internallogger.LogField{"error": err.Error()})
		panic(fmt.Errorf("creating_service %w", err))
	}

	logger.Info("Starting Kafka Consumer")
	setUpKafkaConsumer(configData, srv)
}

func setUpPersistence(ctx context.Context, config *config.AppConf, env string) (*store.DB, error) {
	s := "mongodb+srv"
	if env == "dev" {
		s = "mongodb"
	}

	connection := fmt.Sprintf(
		"%s://%s:%s@%s/<dbname>?retryWrites=true&w=majority",
		s,
		config.DB.Username,
		config.DB.Password,
		config.DB.URL,
	)

	mongoStore, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}

	db, err := store.New(mongoStore)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func setUpKafkaConsumer(config *config.AppConf, srv *service.Service) {
	configMap := kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.Server,
		"group.id":          config.Kafka.GroupID,
	}

	consumer, err := transportkafka.NewConsumer(&configMap, srv)
	if err != nil {
		panic(err)
	}

	topicHandlers := map[string]func(*kafka.Message) error{
		config.Kafka.InsertMovieTopic: consumer.HandleInsertMovie,
	}

	consumer.RegisterTopicHandlers(topicHandlers)

	consumer.StartConsuming()

	select {}
}
