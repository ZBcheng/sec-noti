package handler

import (
	"fmt"
	"net/http"
	"sec-noti/util"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSHandler : websocket接口
func WSHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hanlder init....")
	r.ParseForm()
	conn, _ := upgrader.Upgrade(w, r, nil)
	_, data, err := conn.ReadMessage()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	util.ConnMap[util.MD5(data)] = conn
	fmt.Println(util.ConnMap)
}
