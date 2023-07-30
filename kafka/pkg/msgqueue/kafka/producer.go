package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	pb "github.com/Minsoo-Shin/kafka/api/v1"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"log"
)

type kafkaEventEmitter struct {
	producer sarama.SyncProducer
}

func NewKafkaEventEmitter(client sarama.Client) (msgqueue.EventEmitter, error) {
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}

	emitter := kafkaEventEmitter{
		producer: producer,
	}

	return &emitter, nil
}

func (k *kafkaEventEmitter) Emit(topic pb.Topic, req *pb.Message) error {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic.String(),
		Value: sarama.ByteEncoder(jsonBody),
	}

	log.Printf("published message with topic %s: %v", topic.String(), jsonBody)
	_, _, err = k.producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	return err
}
