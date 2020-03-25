package main

import (
	"net/http"
	"sec-noti/handler"
	"sec-noti/util"
)

// TODO conn回收
func main() {

	go handler.PublishToChannel()
	go writeMessage()
	http.HandleFunc("/noti", handler.WSHandler)
	http.ListenAndServe(":7000", nil)
}

func writeMessage() {
	for {
		msg := <-util.MsgChannel
		for _, conn := range util.ConnMap {
			go conn.WriteMessage(1, []byte(msg))
		}
	}
}
