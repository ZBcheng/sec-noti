package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zbcheng/sec-noti/message"
	"github.com/zbcheng/sec-noti/util"

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
func WSHandler(c *gin.Context) {
	conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
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
			message.ConnMap[resp.Username] = conn

			// 添加新注册用户至postgres映射
			if _, isInMap := message.UserMap[resp.Username]; !isInMap {
				err := message.AddMember(resp.Username)
				if err != nil {
					fmt.Println("Failed to add memeber, err: ", err.Error())
				}
				fmt.Println(message.UserMap)
			} else {
				fmt.Println("member alreay exists in UserMap")
			}
			fmt.Println(message.UserMap)
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
