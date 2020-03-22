package handler

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSHandler : websocket接口
func WSHandler(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil)

	for {

		msg := []byte("helloworld")
		conn.WriteMessage(1, msg)

		if err := conn.WriteMessage(1, msg); err != nil {
			return
		}
	}
}
