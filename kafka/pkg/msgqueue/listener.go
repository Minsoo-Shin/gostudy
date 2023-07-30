package msgqueue

import pb "github.com/Minsoo-Shin/kafka/api/v1"

// EventListener describes an interface for a class that can listen to events.
type EventListener interface {
	Listen() (<-chan pb.Message, chan error, error)
}
