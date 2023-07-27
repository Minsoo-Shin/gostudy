package main

import (
	"flag"
	"github.com/IBM/sarama"
	"github.com/Minsoo-Shin/kafka/event/listener"
	"github.com/Minsoo-Shin/kafka/event/rest"
	"github.com/Minsoo-Shin/kafka/pkg/config"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue/kafka"
)

func main() {
	var eventListener msgqueue.EventListener
	var eventEmitter msgqueue.EventEmitter
	var err error

	confPath := flag.String("conf", "./config/config.json", "flag to set the path to the configuration json file")
	flag.Parse()

	newConfig, _ := config.NewConfig(*confPath)

	switch newConfig.MessageBrokerType {
	case "kafka":
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn := kafka.NewKafkaClient()

		eventListener, err = kafka.NewKafkaEventListener(conn, []int32{})
		if err != nil {
			panic(err)
		}

	default:
		panic("Bad message broker type: " + newConfig.MessageBrokerType)
	}

	//dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	processor := listener.EventProcessor{eventListener}
	go processor.ProcessEvents()
	rest.ServeAPI(newConfig.Listen, eventEmitter)
}
