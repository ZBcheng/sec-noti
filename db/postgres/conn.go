package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

var mutex sync.Mutex
var botID int

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0000"
	dbname   = "security_framework"
)

func init() {
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, _ = sql.Open("postgres", pgInfo)
	// db.SetMaxOpenConns(1000)
	fmt.Println(db)

	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to postgres, err: " + err.Error())
		os.Exit(1)
	}

	botID, err = getBotID()
	if err != nil {
		fmt.Println("Failed to get bot id, err: ", err.Error())
		return
	}

	fmt.Printf("init postgres successfully, bot id: %d\n", botID)
}

// DBConn : 返回pgsql连接
func DBConn() *sql.DB {
	return db
}

// Save2DB : messaeg写入posgres
func Save2DB(message string) {
	mutex.Lock()
	defer mutex.Unlock()
	messageTitle := "来自bot的消息"
	messageContent := message

	sendDate := time.Now().Format("2006-01-02")

	stmt, err := db.Prepare("INSERT INTO message_message (message_title, message_content, message_status, send_time, sender_id) values($1, $2, $3, $4, $5)")
	if err != nil {
		fmt.Println("Failed to prepare, err: ", err.Error())
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(messageTitle, messageContent, '0', sendDate, botID)

	if err != nil {
		fmt.Println("Failed to insert, err: ", err.Error())
		return
	}
}

// getBotID : 获取bot user_id
func getBotID() (botID int, err error) {

	rows, err := db.Query("SELECT ID FROM users_userprofile WHERE username='bot'")
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
