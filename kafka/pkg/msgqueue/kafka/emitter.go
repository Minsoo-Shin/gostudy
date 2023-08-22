package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	pb "github.com/Minsoo-Shin/kafka/api/v1"
	"github.com/Minsoo-Shin/kafka/pkg/config"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"log"
)

type kafkaEventEmitter struct {
	producer sarama.SyncProducer
}

func NewKafkaEventEmitter(conf *config.Config) (msgqueue.EventEmitter, error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	client, err := sarama.NewClient(conf.Kafka.MessageBrokers, config)
	if err != nil {
		log.Fatalf("Failed to load kafka client", err)
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}

	emitter := kafkaEventEmitter{
		producer: producer,
	}

	return &emitter, nil
}

func (k *kafkaEventEmitter) Emit(topic pb.Topic, req msgqueue.Event) error {
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
