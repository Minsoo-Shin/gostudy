package rest

import (
	"encoding/json"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"net/http"
)

type EventHandler struct {
	eventEmitter msgqueue.EventEmitter
}

type eventRequest struct {
	UserID int    `json:"userID"`
	Name   string `json:"Name"`
}

func (r eventRequest) EventName() string {
	return r.Name
}

type eventResponse struct {
	Name string `json:"name"`
}

func (e *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request eventRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = e.eventEmitter.Emit(request)
	if err != nil {
		http.Error(w, "err return", http.StatusInternalServerError)
		return
	}

	response := eventResponse{request.Name}
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(response)
}
