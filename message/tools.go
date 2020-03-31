package message

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	rd "sec-noti/cache/redis"
	pg "sec-noti/db/postgres"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

var rdClient *redis.Client
var pgConn *sql.DB

var mutex sync.Mutex

var botID int
var userIDSet []int

// ConnMap : websocket连接映射
var ConnMap = make(map[string]*websocket.Conn)

// MsgChannel : redis消息存储channel
var MsgChannel = make(chan string, 10)

func init() {
	var err error

	rdClient = rd.RDConn()
	pgConn = pg.DBConn()

	if botID, err = getBotID(); err != nil {
		fmt.Println("Failed to get bot id, err: " + err.Error())
		os.Exit(1)
	}

	if userIDSet, err = getUserIDSet(); err != nil {
		fmt.Println("Failed to get user id set, err: " + err.Error())
		os.Exit(1)
	}

}

// Publish2Channel : 订阅redis频道并发布到util.MsgChannel
func Publish2Channel(rdChannel string) {
	fmt.Println("subscribing channel <" + rdChannel + ">")
	pubsub := rdClient.Subscribe(rdChannel) // 订阅redis频道
	if _, err := pubsub.Receive(); err != nil {
		return
	}
	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println("channel <" + rdChannel + "> published: " + msg.Payload)
		MsgChannel <- msg.Payload // 发送消息到 util.MsgChannel
		save2DB(msg.Payload)
	}
}

// WriteMessage :读取MsgChannel中的消息并发送到前端
func WriteMessage() {
	for {
		msg := <-MsgChannel
		for _, conn := range ConnMap {
			go conn.WriteMessage(1, []byte(msg))
		}
	}
}
