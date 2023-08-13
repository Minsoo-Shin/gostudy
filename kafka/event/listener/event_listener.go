package listener

import (
	"errors"
	"fmt"
	pb "github.com/Minsoo-Shin/kafka/api/v1"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type EventProcessor struct {
	Topic         pb.Topic
	EventListener msgqueue.EventListener
	// DB
	mongodb *mongo.Database
}

func (p *EventProcessor) ProcessEvents() {
	received, errors, err := p.EventListener.Listen()
	if err != nil {
		panic(err)
	}

	var retryMsg = make(chan pb.Message)
	for {
		select {
		case evt := <-received:
			time.Sleep(time.Second * 1)
			fmt.Println("sendFCM start")
			go p.sendFCM(evt, errors, retryMsg)
			fmt.Println("sendFCM end")
		case err = <-errors:
			// save error in DB
			fmt.Println("error channel start")
			fmt.Println("error print!!!", err)
			fmt.Println("error channel end")
		case event := <-retryMsg:
			fmt.Printf("retry: %v", event.GetMsg())
			//p.mongodb.Collection("myCollection").
			//	InsertOne(context.Background(), bson.M{"errorEvent": event.GetMsg()})
		}
	}
}

func (p *EventProcessor) sendFCM(event pb.Message, errChan chan error, retryEvent chan pb.Message) {
	var err error
	// FCM Message
	fmt.Printf("fcm service: (%v) to: %v\n", event.GetMsg(), event.GetFcmToken())
	err = errors.New("new Error")
	if err != nil {
		retryEvent <- event
		errChan <- err
	}
}
