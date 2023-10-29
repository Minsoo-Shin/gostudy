package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gostudy/confluent-go/pkg/config"
	"gostudy/confluent-go/pkg/msgqueue"
	"log"
)

type eventEmitter struct {
	producer *kafka.Producer
}

func NewKafkaEventEmitter(cfg *config.Config) msgqueue.Emitter {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.BootstrapServers,
	})
	if err != nil {
		log.Fatalf("producer failed to create: %v", err)
	}
	//defer p.Close()

	return eventEmitter{
		producer: p,
	}
}

func (e eventEmitter) Close() {
	e.producer.Close()
}

func (e eventEmitter) Emit(messages []msgqueue.Event) {
	// Delivery report handler for produced messages
	go func() {
		for e := range e.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	for _, msg := range messages {
		msgB, _ := json.Marshal(msg)
		kafkaMessage := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: func(s string) *string { return &s }(msg.EventName())},
			Value:          msgB}
		err := e.producer.Produce(kafkaMessage, nil)
		fmt.Println(err)
	}

	// Wait for message deliveries before shutting down
	e.producer.Flush(1 * 1000)
}
