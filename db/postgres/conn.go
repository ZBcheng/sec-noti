package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB

var mutex sync.Mutex
var botID int

const (
	host = "127.0.0.1"
	port = "db"
	user = "postgres"
	// password = "0000"
	dbname = "moviesite"
)

func init() {

	pgInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, _ = sql.Open("postgres", pgInfo)
	db.SetMaxOpenConns(1000)

	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to postgres, err: " + err.Error())
		os.Exit(1)
	}

}

// DBConn : 返回pgsql连接
func DBConn() *sql.DB {
	return db
}
