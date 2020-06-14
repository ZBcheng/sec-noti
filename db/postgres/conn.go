package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"github.com/zbcheng/sec-noti/conf"
)

var db *sql.DB

func init() {
	var config conf.Config
	confPath := conf.GetConfPath()
	if _, err := toml.DecodeFile(confPath, &config); err != nil {
		fmt.Println(err)
		return
	}

	var pgInfo string

	if config.PgConf.Password == "" {
		pgInfo = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
			config.PgConf.Host, config.PgConf.Port, config.PgConf.User, config.PgConf.DBName)
	} else {
		pgInfo = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.PgConf.Host, config.PgConf.Port, config.PgConf.User, config.PgConf.DBName, config.PgConf.Password)
	}

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
