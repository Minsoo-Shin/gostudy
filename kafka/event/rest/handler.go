package rest

import (
	"encoding/json"
	"github.com/Minsoo-Shin/kafka/domain/contracts"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"log"
	"net/http"
)

type EventHandler struct {
	eventEmitter msgqueue.EventEmitter
}

type eventResponse struct {
	Name string `json:"name"`
}

func (e *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("http get")
	var request contracts.EventCreatedEvent
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = e.eventEmitter.Emit(&request)
	if err != nil {
		http.Error(w, "err return", http.StatusInternalServerError)
		return
	}

	response := eventResponse{request.Name}
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(response)
}
