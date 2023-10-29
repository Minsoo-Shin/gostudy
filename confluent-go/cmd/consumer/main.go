package main

import (
	"gostudy/confluent-go/pkg/config"
	"gostudy/confluent-go/pkg/msgqueue/kafka"
)

func main() {
	config := config.New()
	// handler 생성 후 NewKafkaEventListener 인자에 주입
	eventListener := kafka.NewKafkaEventListener(config, nil)
	eventListener.Listen()
}
