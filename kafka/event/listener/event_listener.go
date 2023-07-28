package listener

import (
	"fmt"
	"github.com/Minsoo-Shin/kafka/domain/contracts"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"log"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
}

func (p *EventProcessor) ProcessEvents() {
	log.Println("listening or events")

	received, errors, err := p.EventListener.Listen("eventCreated")

	if err != nil {
		panic(err)
	}

	for {
		select {
		case evt := <-received:
			p.handleEvent(evt)
		case err = <-errors:
			fmt.Printf("got error while receiving event: %s\n", err)
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	fmt.Println("!!handle event!!")
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %v created: %v", e.ID, e)
	default:
		log.Printf("unknown event type: %v", e)
	}
}
