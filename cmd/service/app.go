package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/junkd0g/go-api-example/internal/config"
	internalctx "github.com/junkd0g/go-api-example/internal/context"
	internallogger "github.com/junkd0g/go-api-example/internal/logger"
	store "github.com/junkd0g/go-api-example/internal/persistence/mongo"
	"github.com/junkd0g/go-api-example/internal/service"
	transporthttp "github.com/junkd0g/go-api-example/internal/transport/http"
	transportkafka "github.com/junkd0g/go-api-example/internal/transport/kafka"
)

var (
	env        = os.Getenv("env")
	configPath = os.Getenv("configPath")
)

func main() {
	var wg sync.WaitGroup

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

	wg.Add(1)
	go func() {
		logger.Info("Starting Kafka Consumer")
		setUpKafkaConsumer(configData, srv)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		logger.Info("Starting HTTP Server")
		runHttpServer(ctx, configData, srv)
		wg.Done()
	}()

	// Wait for a SIGINT or SIGTERM signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	logger.Info("Shutting down gracefully...")
	wg.Wait()
}

func setUpPersistence(ctx context.Context, config *config.AppConf, env string) (*store.DB, error) {
	s := "mongodb+srv"
	if env == "dev" {
		s = "mongodb"
	}

	connection := fmt.Sprintf(
		"%s://%s:%s@%s",
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

func runHttpServer(ctx context.Context, config *config.AppConf, srv *service.Service) {
	httpServer, err := transporthttp.NewHttpServer(ctx, srv)
	if err != nil {
		panic(err)
	}

	router := httpServer.GetRouter()
	server := &http.Server{
		Addr:              config.Server.Port,
		ReadHeaderTimeout: 10 * time.Second,
		Handler:           router,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(fmt.Errorf("listener_and_serve: %w", err))
	}
}
