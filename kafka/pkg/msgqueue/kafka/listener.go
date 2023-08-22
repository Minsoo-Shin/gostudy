package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	pb "github.com/Minsoo-Shin/kafka/api/v1"
	"github.com/Minsoo-Shin/kafka/pkg/config"
	msgqueue "github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"log"
)

type kafkaEventListener struct {
	topic    pb.Topic
	consumer sarama.Consumer
}

type event struct {
	event pb.Message
	err   error
}

func NewKafkaEventListener(conf *config.Config, topic pb.Topic) (msgqueue.EventListener, error) {
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
		topic:    topic,
		consumer: consumer,
	}

	return listener, nil
}

func (k *kafkaEventListener) Listen() (<-chan msgqueue.Event, chan error, error) {
	results := make(chan pb.Message)
	errors := make(chan error)

	partitions, err := k.consumer.Partitions(k.topic.String())
	if err != nil {
		log.Fatalf("topic")
	}

	for _, partition := range partitions {
		log.Printf("consuming partition %s:%d", k.topic.String(), partition)

		pConsumer, err := k.consumer.ConsumePartition(k.topic.String(), partition, 0)
		if err != nil {
			return nil, nil, err
		}

		go func() {
			for msg := range pConsumer.Messages() {
				log.Printf("received message %v", msg)

				body := pb.Message{}
				err := json.Unmarshal(msg.Value, &body)
				if err != nil {
					errors <- fmt.Errorf("could not JSON-decode message: %v", err)
					continue
				}

				results <- body
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
