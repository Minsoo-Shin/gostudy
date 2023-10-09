package msgqueue

// EventListener describes an interface for a class that can listen to events.
type EventListener interface {
	Listen(topic string) (<-chan Event, chan error, error)
}
