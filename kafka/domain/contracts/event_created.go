package contracts

// EventCreatedEvent is emitted whenever a new event is created
type EventCreatedEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// EventName returns the event's name
func (c *EventCreatedEvent) EventName() string {
	return "eventCreated"
}
