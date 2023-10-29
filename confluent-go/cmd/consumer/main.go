package main

import (
	"gostudy/confluent-go/pkg/config"
	"gostudy/confluent-go/pkg/msgqueue/kafka"
)

func main() {
	cfg := config.New()
	// handler 생성 후 NewKafkaEventListener 인자에 주입
	eventListener := kafka.NewKafkaEventListener(cfg, nil)
	eventListener.Listen()
}
