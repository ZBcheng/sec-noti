package main

import (
	"net/http"
	"sec-noti/handler"
)

// TODO conn回收
func main() {

	go handler.PublishToChannel()
	http.HandleFunc("/noti", handler.WSHandler)
	http.ListenAndServe(":7000", nil)
}
