package main

import (
	"net/http"
	"sec-noti/handler"
	"sec-noti/redishandler"
	"sec-noti/util"
)

func main() {
	go redishandler.Publish2Channel("bot", util.MsgChannel)
	go util.WriteMessage()
	http.HandleFunc("/noti", handler.WSHandler)
	http.ListenAndServe(":7000", nil)
}
