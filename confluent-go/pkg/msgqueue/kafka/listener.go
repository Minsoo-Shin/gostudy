package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gostudy/confluent-go/pkg/config"
	"gostudy/confluent-go/pkg/msgqueue"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type eventListener struct {
	consumer *kafka.Consumer
	sigchan  chan os.Signal
	config   *config.Config
	mapper   msgqueue.EventMapper
	handler  *msgqueue.Processer // interface convention name + er
}

func NewKafkaEventListener(cfg *config.Config, handler *msgqueue.Processer) msgqueue.Listener {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.BootstrapServers,
		"group.id":          cfg.Kafka.GroupID,
		// Start reading from the first message of each assigned
		// partition if there are no previously committed offsets
		// for this group.
		"auto.offset.reset":        cfg.Kafka.AutoOffsetReset,
		"enable.auto.offset.store": false,
	})
	if err != nil {
		log.Fatalf("consumer failed to create: %v", err)
	}

	return eventListener{
		consumer: c,
		sigchan:  sigchan,
		config:   cfg,
		mapper:   msgqueue.NewEventMapper(),
		handler:  handler,
	}
}

func (l eventListener) Listen() {
	_ = l.consumer.SubscribeTopics(l.config.Kafka.Topics, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true
	for run == true {
		select {
		case sig := <-l.sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := l.consumer.Poll(1000)
			switch e := ev.(type) {
			case *kafka.Message:
				if l.handler != nil {
					event, _ := l.mapper.MapEvent(
						func(p *string) string {
							if p != nil {
								return *p
							}
							return ""
						}(e.TopicPartition.Topic),
						e.Value,
					)
					(*l.handler).Process(event)
				}
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))

				// We can store the offsets of the messages manually or let
				// the library do it automatically based on the setting
				// enable.auto.offset.store. Once an offset is stored, the
				// library takes care of periodically committing it to the broker
				// if enable.auto.commit isn't set to false (the default is true).
				// By storing the offsets manually after completely processing
				// each message, we can ensure atleast once processing.
				_, err := c.StoreMessage(e)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%% Error storing offset after message %s:\n",
						e.TopicPartition)
				}

			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}

	}
	fmt.Printf("Closing consumer\n")
}

func (l eventListener) Close() {
	fmt.Printf("Closing consumer\n")
	l.consumer.Close()
}

/*
consumer에 한들러를 주입받아서 처리한다.
topics을 여러개를 받아서 client를 생성할 수 있다.
config
- topic 추가하면 자연스럽게 process할 수 있게 한다.
-
*/
