package rest

import (
	"fmt"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func ServeAPI(listenAddr string, eventEmitter msgqueue.EventEmitter) {
	r := mux.NewRouter()
	r.Methods("get").Path("/event").Handler(&EventHandler{eventEmitter: eventEmitter})

	srv := http.Server{
		Handler:      r,
		Addr:         listenAddr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	fmt.Printf("server start: %v\n", listenAddr)
	log.Fatalf("server failed: %v\n", srv.ListenAndServe().Error())
}
