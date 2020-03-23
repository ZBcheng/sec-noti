package handler

import (
	"net/http"
	"sec-noti/redishandler"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var channel = make(chan string, 10)

// WSHandler : websocket接口
func WSHandler(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil)

	writeMessage(conn)
}

func writeMessage(conn *websocket.Conn) {
	for {
		msg := <-channel
		if err := conn.WriteMessage(1, []byte(msg)); err != nil {
			return
		}
	}
}

// PublishToChannel : 发送消息到channel
func PublishToChannel() {
	redishandler.Subscribe("bot", channel)
}
