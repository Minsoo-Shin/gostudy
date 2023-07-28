package kafka

import (
	"github.com/IBM/sarama"
	"github.com/Minsoo-Shin/kafka/pkg/config"
	"log"
)

func NewKafkaClient(conf *config.Config) sarama.Client {
	kafkaConf := sarama.NewConfig()
	kafkaConf.Producer.Return.Successes = conf.Kafka.ReturnSuccess

	client, err := sarama.NewClient(conf.Kafka.MessageBrokers, kafkaConf)
	if err != nil {
		log.Fatalf("Failed to load kafka client", err)
	}
	return client

}
