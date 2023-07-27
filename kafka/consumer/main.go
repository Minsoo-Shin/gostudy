package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func main() {
	// Kafka 서버 설정
	topic := "my-topic"
	brokerAddress := ":29092"
	config := sarama.NewConfig()

	// Kafka Consumer 생성
	consumer, err := sarama.NewConsumer([]string{brokerAddress}, config)
	if err != nil {
		log.Fatalln("Failed to create Kafka consumer:", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln("Failed to close Kafka consumer:", err)
		}
	}()

	// Kafka Consumer 연결
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalln("Failed to connect to Kafka consumer:", err)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)

			main()
		}
	}()

	// 메시지 수신
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Received message: %s\n", string(msg.Value))
		case <-signals:
			break ConsumerLoop
		}
	}
}
