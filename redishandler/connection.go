package redishandler

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

// Publish2Channel : 订阅redis频道并发布到util.MsgChannel
func Publish2Channel(rdChannel string, chQueue chan string) {
	fmt.Println("subscribing channel <" + rdChannel + ">")
	pubsub := client.Subscribe(rdChannel) // 订阅redis频道
	if _, err := pubsub.Receive(); err != nil {
		return
	}
	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println("channel <" + rdChannel + "> published: " + msg.Payload)
		chQueue <- msg.Payload // 发送消息到 util.MsgChannel
	}
}
