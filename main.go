package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "测试成功",
		})
	})

	//端口
	r.Run("0.0.0.0:80")
}
