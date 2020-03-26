package main

import (
	"net/http"
	"sec-noti/handler"
	"sec-noti/redishandler"
	"sec-noti/util"
)

func main() {
	go redishandler.Publish2Channel("bot", util.MsgChannel) // 发送redis频道消息到util.MsgChannel
	go util.WriteMessage()                                  // 从util.MsgChannel中读取消息并发送到前端
	http.HandleFunc("/noti", handler.WSHandler)
	http.ListenAndServe(":7000", nil)
}
