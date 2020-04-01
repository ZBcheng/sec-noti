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
	host = "db"
	port = 5432
	user = "postgres"
	// password = "0000"
	dbname = "postgres"
)

func init() {
	// pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

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
