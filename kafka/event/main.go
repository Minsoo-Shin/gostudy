package main

import (
	"flag"
	"github.com/Minsoo-Shin/kafka/event/listener"
	"github.com/Minsoo-Shin/kafka/event/rest"
	"github.com/Minsoo-Shin/kafka/pkg/config"
	"github.com/Minsoo-Shin/kafka/pkg/mongo"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue/kafka"
)

func main() {
	var eventListener msgqueue.EventListener
	var eventEmitter msgqueue.EventEmitter
	var err error

	confPath := flag.String("conf", ".pkg/config/config.json", "flag to set the path to the configuration json file")
	flag.Parse()

	conf, _ := config.NewConfig(*confPath)

	eventEmitter, err = kafka.NewKafkaEventEmitter(conf)
	if err != nil {
		panic(err)
	}
	eventListener, err = kafka.NewKafkaEventListener(conf)
	if err != nil {
		panic(err)
	}

	mongodb, err := mongo.New()
	if err != nil {
		panic(err)
	}

	processor := listener.EventProcessor{
		EventListener: eventListener,
		Mongodb:       mongodb,
	}
	go processor.ProcessEvents("eventCreated")

	rest.ServeAPI(conf.Listen, eventEmitter)
}
