package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var RoomId = 0

var Rooms = make(map[*Room]bool, 0)

type Room struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Clients map[string]*Client
}

func createRoom(hub *Hub, w http.ResponseWriter, r *http.Request) {
	var roomRequest Room
	err := json.NewDecoder(r.Body).Decode(&roomRequest)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	bToken := r.Header.Get("Authorization")
	claims := ParseToken(bToken[7:])

	for r := range Rooms {
		if r.Name == roomRequest.Name {
			fmt.Println("find room")
		}
		http.Error(w, "try other room name", http.StatusBadRequest)
		return
	}
	fmt.Println("can't find room so can make room")

	var newRoom = &Room{
		Id:   RoomId,
		Name: roomRequest.Name,
		Clients: map[string]*Client{
			claims["memberId"].(string): {},
		},
	}

	RoomId++

	Rooms[newRoom] = true
	http.ServeFile(w, r, "chat.html")
}

func joinRoom(w http.ResponseWriter, r *http.Request) {

}
