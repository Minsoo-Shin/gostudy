package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gostudy/confluent-go/pkg/config"
	"gostudy/confluent-go/pkg/msgqueue"
	"log"
	"os"
)

type eventListener struct {
	consumer *kafka.Consumer
	config   *config.Config
	handler  *msgqueue.Processer // interface convention name + er
}

func NewKafkaEventListener(cfg *config.Config, handler *msgqueue.Processer) msgqueue.Listener {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.BootstrapServers,
		"group.id":          cfg.Kafka.GroupID,
		"auto.offset.reset": cfg.Kafka.AutoOffsetReset,
	})
	if err != nil {
		log.Fatalf("consumer failed to create: %v", err)
	}
	return eventListener{
		consumer: c,
		config:   cfg,
		handler:  handler,
	}
}

func (l eventListener) Listen() {
	_ = l.consumer.SubscribeTopics(l.config.Kafka.Topics, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true
	for run == true {
		ev := l.consumer.Poll(1000)
		switch e := ev.(type) {
		case *kafka.Message:
			if l.handler != nil {
				(*l.handler).Process(e)
			}
			fmt.Printf("%% Message on %s:\n%s\n",
				e.TopicPartition, string(e.Value))
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}

	l.consumer.Close()
}

/*
consumer에 한들러를 주입받아서 처리한다.
topics을 여러개를 받아서 client를 생성할 수 있다.
config
- topic 추가하면 자연스럽게 process할 수 있게 한다.
-
*/
