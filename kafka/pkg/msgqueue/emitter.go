package msgqueue

import (
	pb "github.com/Minsoo-Shin/kafka/api/v1"
)

type EventEmitter interface {
	Emit(topic pb.Topic, req Event) error
}
