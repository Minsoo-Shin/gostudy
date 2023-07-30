package main

import (
	"flag"
	pb "github.com/Minsoo-Shin/kafka/api/v1"
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

	conf, _ := config.NewConfig(*confPath)

	conn := kafka.NewKafkaClient(conf)

	eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
	if err != nil {
		panic(err)
	}
	eventListener, err = kafka.NewKafkaEventListener(conn, pb.Topic_AclosetNotification, []int32{})
	if err != nil {
		panic(err)
	}

	processor := listener.EventProcessor{EventListener: eventListener}
	go processor.ProcessEvents()

	rest.ServeAPI(conf.Listen, eventEmitter)
}
