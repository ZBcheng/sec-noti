package redishandler

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

// Publish2Channel : 订阅redis频道并发布到 go channel
func Publish2Channel(rdChannel string, chQueue chan string) {
	fmt.Println("subscribing channel <" + rdChannel + ">")
	pubsub := client.Subscribe(rdChannel)
	if _, err := pubsub.Receive(); err != nil {
		return
	}
	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Payload)
		chQueue <- msg.Payload
	}
}
