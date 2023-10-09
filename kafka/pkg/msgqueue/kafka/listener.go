package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/Minsoo-Shin/kafka/pkg/config"
	msgqueue "github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"log"
)

type kafkaEventListener struct {
	mapper   msgqueue.EventMapper
	consumer sarama.Consumer
}

type event struct {
	event msgqueue.Event
	err   error
}

func NewKafkaEventListener(conf *config.Config) (msgqueue.EventListener, error) {
	config := sarama.NewConfig()

	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	client, err := sarama.NewClient(conf.Kafka.MessageBrokers, config)
	if err != nil {
		log.Fatalf("Failed to load kafka client", err)
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return nil, err
	}

	listener := &kafkaEventListener{
		mapper:   msgqueue.NewEventMapper(),
		consumer: consumer,
	}

	return listener, nil
}

func (k *kafkaEventListener) Listen(topic string) (<-chan msgqueue.Event, chan error, error) {
	results := make(chan msgqueue.Event)
	errors := make(chan error)

	partitions, err := k.consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("topic")
	}

	for _, partition := range partitions {
		log.Printf("consuming partition %s:%d", partition)

		pConsumer, err := k.consumer.ConsumePartition(topic, partition, 0)
		if err != nil {
			return nil, nil, err
		}

		go func() {
			for msg := range pConsumer.Messages() {
				log.Printf("received message %v", msg)

				body := &MsgEnvelope{}
				err := json.Unmarshal(msg.Value, &body)
				if err != nil {
					errors <- fmt.Errorf("could not JSON-decode message: %v", err)
					continue
				}

				event, err := k.mapper.MapEvent(body.EventName, body.Message)
				if err != nil {
					errors <- fmt.Errorf("could not map message: %v", err)
					continue
				}

				results <- event
			}
		}()

		go func() {
			for err := range pConsumer.Errors() {
				errors <- err
			}
		}()
	}

	return results, errors, nil
}
