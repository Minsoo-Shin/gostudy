package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	ListenDefault              = "127.0.0.1:3000"
	MessageBrokerTypeDefault   = "kafka"
	KafkaMessageBrokersDefault = []string{"localhost:29092"}
)

type Config struct {
	Listen              string   `json:"listen"`
	MessageBrokerType   string   `json:"message_broker_type"`
	KafkaMessageBrokers []string `json:"kafka_message_brokers"`
}

func NewConfig(filename string) (*Config, error) {
	conf := &Config{
		ListenDefault,
		MessageBrokerTypeDefault,
		KafkaMessageBrokersDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	if err = json.NewDecoder(file).Decode(&conf); err != nil {
		log.Fatalf("decode err: %v", err)
	}

	if v := os.Getenv("LISTEN"); v != "" {
		conf.Listen = v
	}

	if v := os.Getenv("KAFKA_BROKER_URLS"); v != "" {
		conf.MessageBrokerType = "kafka"
		conf.KafkaMessageBrokers = strings.Split(v, ",")
	}

	return conf, nil
}
