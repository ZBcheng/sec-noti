package main

import (
	"net/http"
	"sec-noti/handler"
)

func main() {

	go handler.PublishToChannel()
	http.HandleFunc("/noti", handler.WSHandler)
	http.ListenAndServe(":7000", nil)
}
