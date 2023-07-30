package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	pb "github.com/Minsoo-Shin/kafka/api/v1"
	msgqueue "github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"log"
)

type kafkaEventListener struct {
	topic      pb.Topic
	consumer   sarama.Consumer
	partitions []int32
}

func NewKafkaEventListener(client sarama.Client, topic pb.Topic, partitions []int32) (msgqueue.EventListener, error) {
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return nil, err
	}

	listener := &kafkaEventListener{
		topic:      topic,
		consumer:   consumer,
		partitions: partitions,
	}

	return listener, nil
}

func (k *kafkaEventListener) Listen() (<-chan pb.Message, chan error, error) {
	var err error

	results := make(chan pb.Message)
	errors := make(chan error)

	partitions := k.partitions
	if len(partitions) == 0 {
		partitions, err = k.consumer.Partitions(k.topic.String())
		if err != nil {
			return nil, nil, err
		}
	}

	log.Printf("topic %s has partitions: %v", k.topic.String(), partitions)

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
