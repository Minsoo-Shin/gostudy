package msgqueue

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type Emitter interface {
	Emit(messages []*kafka.Message)
	Close()
}
