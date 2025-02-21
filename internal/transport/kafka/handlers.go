package transportkafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	internalctx "github.com/IordanisCodeExamples/go-api-example/internal/context"
	transport "github.com/IordanisCodeExamples/go-api-example/internal/transport"
)

// HandleInsertMovie kafka handler to insert movie operation
func (c Consumer) HandleInsertMovie(msg *kafka.Message) error {
	logger, _ := internalctx.GetLoggerFromContext(c.ctx)
	logger.Info("getMovie operation started")

	var movie transport.Movie
	if err := json.Unmarshal(msg.Value, &movie); err != nil {
		logger.Error(err.Error())
		return err
	}

	if err := c.Service.IngestMovie(c.ctx, movie); err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("getMovie operation ended")
	return nil
}
