module sec-noti

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gorilla/websocket v1.4.2
	github.com/lib/pq v1.3.0
	github.com/zbcheng/sec-noti v0.0.0-20200602171035-44e96326db0f
)

replace github.com/zbcheng/sec-noti => ./
