package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sec-noti/util"

	"github.com/gorilla/websocket"
)

// WebSocketResp : websocket返回信息
type WebSocketResp struct {
	Username  string `json:"username"`
	Operation int    `json:"operation"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSHandler : websocket接口
func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	for {
		_, rawData, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data := string(rawData)
		resp := WebSocketResp{}

		if err := json.Unmarshal([]byte(data), &resp); err != nil {
			return
		}

		if resp.Operation == 0 {
			// Operation == 0 means user is connecting
			util.ConnMap[util.MD5(resp.Username)] = conn
			fmt.Println("user: " + resp.Username + " connected!")
		} else if resp.Operation == 1 {
			// Operation == 1 means user is logging out
			delete(util.ConnMap, util.MD5(resp.Username))
			fmt.Println("user: " + resp.Username + " disconnected!")
		} else {
			panic("Unknown Operation!")
		}

		fmt.Println(util.ConnMap)
	}

}
