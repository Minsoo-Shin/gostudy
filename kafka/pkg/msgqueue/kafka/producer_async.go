package kafka

import (
	"github.com/IBM/sarama"
)

type kafkaAsyncEventEmitter struct {
	producer sarama.AsyncProducer
}

//
//func NewKafkaAsyncEventEmitter(client sarama.Client) (msgqueue.EventEmitter, error) {
//	producer, err := sarama.NewAsyncProducerFromClient(client)
//	if err != nil {
//		return nil, err
//	}
//
//	emitter := kafkaAsyncEventEmitter{
//		producer: producer,
//	}
//
//	return &emitter, nil
//}
//
//func (k *kafkaAsyncEventEmitter) Emit(topic pb.Topic, req *pb.Message, signals chan os.Signal) error {
//	jsonBody, err := json.Marshal(req)
//	if err != nil {
//		return err
//	}
//
//	msg := &sarama.ProducerMessage{
//		Topic: topic.String(),
//		Value: sarama.ByteEncoder(jsonBody),
//	}
//
//	log.Printf("published message with topic %s: %v", topic.String(), jsonBody)
//	for {
//		select {
//		case k.producer.Input() <- msg:
//			log.Println("New Message produced")
//		case <-signals:
//			k.producer.AsyncClose() // Trigger a shutdown of the producer.
//			return nil
//		}
//	}
//}
