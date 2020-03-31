package message

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	rd "sec-noti/cache/redis"
	pg "sec-noti/db/postgres"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

var rdClient *redis.Client
var pgConn *sql.DB

var mutex sync.Mutex

var botID int

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

// Save2DB : messaeg写入posgres
func Save2DB(message string) (err error) {
	mutex.Lock()
	defer mutex.Unlock()
	messageTitle := "来自bot的消息"
	messageContent := message

	sendDate := time.Now().Format("2006-01-02")

	stmt, err := pgConn.Prepare("INSERT INTO message_message (message_title, message_content, message_status, send_time, sender_id) values($1, $2, $3, $4, $5)")
	if err != nil {
		fmt.Println("Failed to prepare, err: ", err.Error())
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(messageTitle, messageContent, '0', sendDate, botID)

	if err != nil {
		fmt.Println("Failed to insert, err: ", err.Error())
		return err
	}

	return nil
}

// getBotID : 获取bot user_id
func getBotID() (botID int, err error) {

	rows, err := pgConn.Query("SELECT ID FROM users_userprofile WHERE username='bot'")
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		err = rows.Scan(&botID)
		if err != nil {
			return 0, err
		}
	}

	return botID, nil
}
