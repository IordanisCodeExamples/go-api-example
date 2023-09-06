package transportkafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (c Consumer) HandleInsertMovie(msg *kafka.Message) error {
	fmt.Println("HandleInsertMovie")
	return nil
}
