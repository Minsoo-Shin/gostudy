package rest

import (
	"encoding/json"
	pb "github.com/Minsoo-Shin/kafka/api/v1"
	"github.com/Minsoo-Shin/kafka/pkg/msgqueue"
	"net/http"
)

type EventHandler struct {
	eventEmitter msgqueue.EventEmitter
}

type eventResponse struct {
	Name string `json:"name"`
}

func (e *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request pb.Message
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = e.eventEmitter.Emit(pb.Topic_AclosetNotification, &request)
	if err != nil {
		http.Error(w, "err return", http.StatusInternalServerError)
		return
	}

	response := eventResponse{request.GetMsg()}
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(response)
}
