package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
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
	config.Producer.Retry.Max = 5                    // Retry up to 5 times to produce the message
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

type MsgEnvelope struct {
	EventName string      `json:"topic"`
	Message   interface{} `json:"message"`
}

func (k *kafkaEventEmitter) Emit(req msgqueue.Event) error {
	envelope := MsgEnvelope{
		EventName: req.EventName(),
		Message:   req,
	}
	jsonBody, err := json.Marshal(envelope)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: req.EventName(),
		Value: sarama.ByteEncoder(jsonBody),
	}

	log.Printf("published message with topic %s", req.EventName())
	_, _, err = k.producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	return err
}
