package handler

import (
	"fmt"
	"net/http"
	"sec-noti/redishandler"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var channel = make(chan string, 10) // redis消息存储channel
var connPool []*websocket.Conn      // websocket连接池

// WSHandler : websocket接口
func WSHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hanlder init....")
	conn, _ := upgrader.Upgrade(w, r, nil)
	connPool = append(connPool, conn)
	fmt.Println(connPool)

	writeMessage()
}

// writeMessage : 向前端返回信息
func writeMessage() {
	for {
		msg := <-channel
		for _, conn := range connPool {
			go conn.WriteMessage(1, []byte(msg))
		}
	}
}

// PublishToChannel : 发送消息到channel
func PublishToChannel() {
	redishandler.Subscribe("bot", channel)
}
