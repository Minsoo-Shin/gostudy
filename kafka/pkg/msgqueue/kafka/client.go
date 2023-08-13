package kafka

import (
	"github.com/IBM/sarama"
	"github.com/Minsoo-Shin/kafka/pkg/config"
	"log"
)

func NewKafkaClient(conf *config.Config) sarama.Client {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	client, err := sarama.NewClient(conf.Kafka.MessageBrokers, config)
	if err != nil {
		log.Fatalf("Failed to load kafka client", err)
	}
	return client
}
