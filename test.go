package main

import "sec-noti/util"

func main() {
	testMsg := []byte("world")
	util.MD5(testMsg)
}