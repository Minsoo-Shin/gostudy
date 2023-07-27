package kafka

import (
	"github.com/IBM/sarama"
	"log"
	"os"
	"strings"
)

func NewKafkaClient() sarama.Client {
	brokerList := os.Getenv("KAFKA_BROKERS")
	if brokerList == "" {
		brokerList = "localhost:29092"
	}

	brokers := strings.Split(brokerList, ",")
	config := sarama.NewConfig()
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		log.Fatalf("Failed to load kafka client", err)
	}
	return client

}
