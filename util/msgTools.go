package util

import "github.com/gorilla/websocket"

// ConnMap : websocket连接映射
var ConnMap = make(map[string]*websocket.Conn)

// MsgChannel : redis消息存储channel
var MsgChannel = make(chan string, 10)

// WriteMessage :读取MsgChannel中的消息
func WriteMessage() {
	for {
		msg := <-MsgChannel
		for _, conn := range ConnMap {
			go conn.WriteMessage(1, []byte(msg))
		}
	}
}
