package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "测试成功",
		})
	})

	r.LoadHTMLGlob("view/*")
	r.Static("/css", "./css")
	r.Static("/js", "./js")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "hello Go",
		})
	})

	//端口
	r.Run("0.0.0.0:80")
}
