package main

import "sec-noti/redishandler"

func main() {
	channel := make(chan string, 10)
	redishandler.Subscribe("bot", channel)
}
