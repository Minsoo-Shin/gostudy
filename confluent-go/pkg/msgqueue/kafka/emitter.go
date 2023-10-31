package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gostudy/confluent-go/pkg/config"
	"gostudy/confluent-go/pkg/msgqueue"
	"log"
	"time"
)

type eventEmitter struct {
	producer *kafka.Producer
}

func NewKafkaEventEmitter(cfg *config.Config) msgqueue.Emitter {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  cfg.Kafka.BootstrapServers,
		"enable.idempotence": true,
	})
	if err != nil {
		log.Fatalf("producer failed to create: %v", err)
	}

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
			case kafka.Error:
				e := ev
				if e.IsFatal() {
					fmt.Printf("FATAL ERROR: %v: terminating\n", e)
				} else {
					fmt.Printf("Error: %v\n", e)
				}

			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	sendidx := 0
	for sendidx < len(messages) {
		msg := messages[sendidx]
		msgB, _ := json.Marshal(msg)
		kafkaMessage := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     func(s string) *string { return &s }(msg.EventName()),
				Partition: kafka.PartitionAny,
			},
			Value: msgB,
		}
		err := e.producer.Produce(kafkaMessage, nil)

		if err != nil {
			if err.(kafka.Error).Code() == kafka.ErrQueueFull {
				// Producer queue is full, wait 1s for messages
				// to be delivered then try again.
				time.Sleep(time.Second)
				continue
			}
			// 메세지 전송 실패 -> 처리?
			fmt.Printf("Failed to produce message: %v\n", err)
		}
		sendidx++
	}

	// Wait for message deliveries before shutting down
	for e.producer.Flush(10000) > 0 {
		fmt.Print("Still waiting to flush outstanding messages\n")
	}
}
