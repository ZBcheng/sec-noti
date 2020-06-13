package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
	"github.com/BurntSushi/toml"
)

var db *sql.DB

var mutex sync.Mutex
var botID int

type pgConf struct {
	host string
	port int
	user string
	dbname string
	password string
}

// const (
// 	host   = "db"
// 	port   = 5432
// 	user   = "postgres"
// 	dbname = "postgres"
// )

func init() {
	var pg pgConf
	confPath := "./conf_aliyun.toml"
	if _, err := toml.DecodeFile(confPath, &pg); err != nil {
		fmt.Println(err)
		return
	}
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		pg.host, pg.port, pg.user, pg.dbname, pg.password)
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
