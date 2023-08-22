package contracts

import pb "github.com/Minsoo-Shin/kafka/api/v1"

type EventNew struct {
	pb.Message
}

// EventName returns the event's name
func (c *EventNew) EventName() string {
	return "eventCreated"
}
