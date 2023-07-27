package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

func main() {
	// Kafka 서버 설정
	topic := "my-topic"
	brokerAddress := ":29092"

	// Kafka Producer 구성
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Return.Successes = true

	// Kafka Producer 생성
	producer, err := sarama.NewSyncProducer([]string{brokerAddress}, producerConfig)
	if err != nil {
		log.Fatalln("Failed to create Kafka producer:", err)
	}

	// 메시지 전송
	i := 0
	for {
		message := fmt.Sprintf("i'm %v", i)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatalln("Failed to send message:", err)
		}
		fmt.Printf("Message sent successfully! Partition: %d, Offset: %d\n", partition, offset)
		time.Sleep(time.Second * 3)
		i++
	}
}
