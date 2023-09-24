package main

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var (
	MemberDB = []Member{}
)

type Member struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

type MemberWithToken struct {
	Member
	Token string `json:"token"`
}

func init() {
	MemberDB = append(MemberDB, Member{
		Id:       1,
		Username: "test1",
		Password: "123",
	}, Member{
		Id:       2,
		Username: "test2",
		Password: "123",
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "login.html")
		return
	case http.MethodPost:
		var member Member
		if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, m := range MemberDB {
			if m.Username == member.Username && m.Password == member.Password {
				m.Password = ""

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"memberId": m.Id,
					"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
				})

				//Sign and get the complete encoded token as a string using the secret
				tokenString, _ := token.SignedString([]byte("secret"))

				err := json.NewEncoder(w).Encode(MemberWithToken{
					Member: m,
					Token:  tokenString,
				})
				if err != nil {
					http.Error(w, "internal error", http.StatusInternalServerError)
				}
				return
			}
		}
		http.Error(w, "check username and password", http.StatusUnauthorized)
	default:
		return
	}
}
