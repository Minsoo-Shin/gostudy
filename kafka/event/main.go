package main

import (
	"flag"
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
	eventListener, err = kafka.NewKafkaEventListener(conn, []int32{})
	if err != nil {
		panic(err)
	}

	//dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	processor := listener.EventProcessor{EventListener: eventListener}
	go processor.ProcessEvents()
	rest.ServeAPI(conf.Listen, eventEmitter)
}
