package handler

import (
	"fmt"
	"net/http"
	"sec-noti/redishandler"
	"sec-noti/util"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// var channel = make(chan string, 10) // redis消息存储channel
// var connMap = util.ConnMap          // websocket连接池

// WSHandler : websocket接口
func WSHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hanlder init....")
	r.ParseForm()
	conn, _ := upgrader.Upgrade(w, r, nil)
	_, data, err := conn.ReadMessage()
	if err != nil {
		fmt.Println(err.Error())
	}
	util.ConnMap[util.MD5(data)] = conn
	fmt.Println(util.ConnMap)
	// writeMessage()
}

// writeMessage : 向前端返回信息

// PublishToChannel : 发送消息到channel
func PublishToChannel() {
	redishandler.Subscribe("bot", util.MsgChannel)
}
