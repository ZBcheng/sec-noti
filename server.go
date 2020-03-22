package main

import (
	"net/http"
	"sec-noti/handler"
)

func main() {
	http.HandleFunc("/ws", handler.WSHandler)
	http.ListenAndServe("0.0.0.0:7000", nil)
}
