package kafka

import (
	"github.com/IBM/sarama"
	"github.com/Minsoo-Shin/kafka/pkg/config"
	"log"
	"time"
)

func NewKafkaClient(conf *config.Config) sarama.Client {
	kafkaConf := sarama.NewConfig()

	if conf.Kafka.IsAsync {
		// 비동기
		kafkaConf.Producer.RequiredAcks = sarama.NoResponse
		kafkaConf.Producer.Compression = sarama.CompressionSnappy
		kafkaConf.Producer.Flush.Frequency = 1 * time.Millisecond
	} else {
		// 동기
		kafkaConf.Producer.Return.Successes = conf.Kafka.ReturnSuccess
	}

	client, err := sarama.NewClient(conf.Kafka.MessageBrokers, kafkaConf)
	if err != nil {
		log.Fatalf("Failed to load kafka client", err)
	}
	return client
}
