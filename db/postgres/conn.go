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

	pgInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.PgConf.Host, config.PgConf.Port, config.PgConf.User, config.PgConf.DBName, config.PgConf.Password)
	db, _ = sql.Open("postgres", pgInfo)
	db.SetMaxOpenConns(1000)

	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to postgres, err: " + err.Error())
		os.Exit(1)
	}

}

func loadConfig() {
	var conf conf.Config
	confPath := "conf/conf.toml"
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(conf.PgConf.Host)
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		conf.PgConf.Host, conf.PgConf.Port, conf.PgConf.User, conf.PgConf.DBName, conf.PgConf.Password)
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
