package main

import (
	"gostudy/confluent-go/domain/contracts"
	"gostudy/confluent-go/pkg/config"
	"gostudy/confluent-go/pkg/msgqueue"
	"gostudy/confluent-go/pkg/msgqueue/kafka"
)

func main() {

	cfg := config.New()
	/*
		handler 생성 후
		NewKafkaEventListener 인자에 주입
	*/
	eventListener := kafka.NewKafkaEventEmitter(cfg)
	defer eventListener.Close()

	for _, v := range []string{"test", "test2", "test3"} {
		event := &contracts.EventCreatedEvent{
			EventID:      1,
			EventMessage: v,
		}
		eventListener.Emit([]msgqueue.Event{event})
	}

}
