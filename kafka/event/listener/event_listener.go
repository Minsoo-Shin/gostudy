package listener

import (
	"fmt"
	pb "github.com/Minsoo-Shin/kafka/api/v1"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"log"
)

type EventProcessor struct {
	Topic         pb.Topic
	EventListener msgqueue.EventListener
	// DB
}

func (p *EventProcessor) ProcessEvents() {
	received, errors, err := p.EventListener.Listen()
	if err != nil {
		panic(err)
	}

	for {
		select {
		case evt := <-received:
			p.sendMessage(evt, errors)
		case err = <-errors:
			fmt.Printf("got error while receiving event: %s\n", err)
			// save error in DB
		}
	}
}

func (p *EventProcessor) sendMessage(event pb.Message, errChan chan error) {
	var err error
	// FCM Message
	log.Printf("msg: %v\n to: %v\n", event.GetMsg(), event.GetFcmToken())
	if err != nil {
		errChan <- err
	}
}
