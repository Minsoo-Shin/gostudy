package listener

import (
	"context"
	"fmt"
	"github.com/Minsoo-Shin/kafka/domain/contracts"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
	// DB
	Mongodb *mongo.Database
}

func (p *EventProcessor) ProcessEvents(topic string) {
	received, errors, err := p.EventListener.Listen(topic)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case evt := <-received:
			time.Sleep(time.Second * 1)
			fmt.Println("sendFCM start")
			p.handleEvent(evt)
			fmt.Println("sendFCM end")
		case err = <-errors:
			// save error in DB
			fmt.Println("error channel start")
			fmt.Println("error print!!!", err)
			fmt.Println("error channel end")
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %v created: %v", e.EventID, e)
		_, err := p.Mongodb.Collection("event").InsertOne(context.Background(), e)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
		return
	default:
		log.Printf("unknown event type: %T", e)
	}
}
