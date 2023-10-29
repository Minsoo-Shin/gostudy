package msgqueue

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type Processer interface {
	Process(event *kafka.Message)
}
