package util

import (
	"crypto/md5"

	"github.com/gorilla/websocket"
)

var ConnMap = make(map[string]*websocket.Conn)
var MsgChannel = make(chan string, 10)

// MD5 : md5算法加密
func MD5(rawData []byte) string {
	enc := md5.New()
	enc.Write(rawData)
	result := enc.Sum([]byte(""))
	return string(result)
}
