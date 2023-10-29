package contracts

type EventCreatedEvent struct {
	EventID      int    `json:"eventID"`
	EventMessage string `json:"eventMessage"`
}

// EventName returns the event's name
func (c *EventCreatedEvent) EventName() string {
	return "eventCreated"
}
