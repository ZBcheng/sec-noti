package redis

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis"
	"github.com/zbcheng/sec-noti/conf"
)

var client *redis.Client

// init : 初始化连接
func init() {
	var config conf.Config
	confPath := conf.GetConfPath()
	if _, err := toml.DecodeFile(confPath, &config); err != nil {
		fmt.Println(err)
		return
	}
	addr := fmt.Sprintf("%s:%s", config.RdConf.Host, config.RdConf.Port)
	client = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

// RDConn : 返回redis客户端
func RDConn() *redis.Client {
	return client
}
