package main

import (
	"net/http"
	"sec-noti/handler"
	"sec-noti/message"
)

func main() {
	go message.Publish2Channel("bot") // 发送redis频道消息到util.MsgChannel
	go message.WriteMessage()         // 从util.MsgChannel中读取消息并发送到前端
	http.HandleFunc("/noti", handler.WSHandler)
	http.ListenAndServe(":7000", nil)
}
