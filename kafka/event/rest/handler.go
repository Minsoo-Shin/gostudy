package rest

import (
	"encoding/json"
	"github.com/Minsoo-Shin/kafka/domain/contracts"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"net/http"
)

type EventHandler struct {
	eventEmitter msgqueue.EventEmitter
}

type eventRequest struct {
	EventName    string `json:"eventName"`
	EventMessage string `json:"eventMessage"`
}

type eventResponse struct {
	Status string `json:"status"`
}

func (e *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request eventRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = e.eventEmitter.Emit(&contracts.EventCreatedEvent{
		EventID:      1,
		EventMessage: request.EventMessage,
	})
	if err != nil {
		http.Error(w, "err return", http.StatusInternalServerError)
		return
	}

	response := eventResponse{Status: "success"}
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(response)
}
