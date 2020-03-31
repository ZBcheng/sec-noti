package redis

import "github.com/go-redis/redis"

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

// RDConn : 返回redis客户端
func RDConn() *redis.Client {
	return client
}
