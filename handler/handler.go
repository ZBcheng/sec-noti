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

	c.JSON(http.StatusOK, gin.H{
		"msg":  "",
		"data": usernameSet,
		"err":  "",
	})
}
