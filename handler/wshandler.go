package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sec-noti/message"
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
		_, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		resp := WebSocketResp{}

		if err := json.Unmarshal(data, &resp); err != nil {
			return
		}

		if resp.Operation == 0 {
			// Operation == 0 means user is connecting
			message.ConnMap[util.MD5(resp.Username)] = conn

			if _, isInMap := message.UserMap[resp.Username]; !isInMap {
				err := message.AddMember(resp.Username)
				if err != nil {
					fmt.Println("Failed to add memeber, err: ", err.Error())
				}
				fmt.Println(message.UserMap)
			}
			fmt.Println("user: " + resp.Username + " connected!")
		} else if resp.Operation == 1 {
			// Operation == 1 means user is logging out
			delete(message.ConnMap, util.MD5(resp.Username))
			fmt.Println("user: " + resp.Username + " disconnected!")
		} else {
			panic("Unknown Operation!")
		}

		fmt.Println(message.ConnMap)
	}

}
