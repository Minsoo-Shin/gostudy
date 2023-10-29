package main

import (
	kafkalib "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gostudy/confluent-go/pkg/config"
	"gostudy/confluent-go/pkg/msgqueue/kafka"
)

func main() {

	cfg := config.New()
	/*
		handler 생성 후
		NewKafkaEventListener 인자에 주입
	*/
	eventListener := kafka.NewKafkaEventEmitter(cfg)
	defer eventListener.Close()

	for _, v := range []string{"test", "test2", "test3"} {
		eventListener.Emit([]*kafkalib.Message{
			{
				TopicPartition: kafkalib.TopicPartition{
					Topic: func(s string) *string { return &s }("myTopic"),
				},
				Value: []byte(v),
			},
		})
	}

}
