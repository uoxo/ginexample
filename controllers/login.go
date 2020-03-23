package controllers

import (
	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (this LoginController) Login(c *gin.Context) {
	tip := "登入成功"
	c.JSON(200, gin.H{"code": "101", "tip": tip, "data": ""})
	c.Abort()
}
