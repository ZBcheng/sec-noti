package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zbcheng/sec-noti/message"
)



// ConnectedUsers : 查看连接用户
func ConnectedUsers(c *gin.Context) {
	usernameSet := make([]string, 0)
	for username, _ := range message.ConnMap {
		usernameSet = append(usernameSet, username)
	}

	num := len(usernameSet)

	c.JSON(http.StatusOK, gin.H{
		"data": usernameSet,
		"number": num,
		"msg":  "",
		"err":  "",
	})
}


func IsAlive(c *gin.Context) {}