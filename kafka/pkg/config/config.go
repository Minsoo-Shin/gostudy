package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

const (
	ListenDefault              = "127.0.0.1:3000"
	KafkaMessageBrokersDefault = "localhost:29092"
)

type Config struct {
	Listen string    `json:"listen"`
	Kafka  KafkaConf `json:"kafka"`
}

type KafkaConf struct {
	MessageBrokers []string `json:"messageBrokers"`
	IsAsync        bool     `json:"isAsync"`
	ReturnSuccess  bool     `json:"returnSuccess"`
}

func NewConfig(filename string) (*Config, error) {
	// set default conf
	conf := &Config{
		Listen: ListenDefault,
		Kafka: KafkaConf{
			MessageBrokers: []string{KafkaMessageBrokersDefault},
			ReturnSuccess:  true,
		},
	}
	// via file
	file, err := os.Open(filename)
	if err != nil {
		return conf, err
	}

	if err = json.NewDecoder(file).Decode(&conf); err != nil {
		log.Fatalf("err decode: %v", err)
	}

	if v := os.Getenv("LISTEN"); v != "" {
		conf.Listen = v
	}

	if v := os.Getenv("KAFKA_BROKERS"); v != "" {
		conf.Kafka.MessageBrokers = strings.Split(v, ",")
	}

	return conf, nil
}
