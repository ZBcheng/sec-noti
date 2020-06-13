module sec-noti

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/aerospike/aerospike-client-go v3.0.0+incompatible
	github.com/arstd/log v0.0.0-20200414075513-0888823dd60f
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gorilla/websocket v1.4.2
	github.com/lib/pq v1.3.0
	github.com/zbcheng/sec-noti v0.0.0-20200602171035-44e96326db0f
)

replace github.com/zbcheng/sec-noti => ./
