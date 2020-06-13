package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zbcheng/sec-noti/handler"
	"github.com/zbcheng/sec-noti/message"
	"github.com/zbcheng/sec-noti/util"
)

func main() {
	go message.Publish2Channel("bot") // 发送redis频道消息到message.MsgChannel
	go message.WriteMessage()         // 从message.MsgChannel中读取消息并发送到前端

	router := gin.Default()
	router.Use(util.Cors())

	router.GET("/noti/users", handler.ConnectedUsers)
	router.GET("/noti", handler.WSHandler)

	router.Run(":7000")
}
