package handler

import "github.com/gin-gonic/gin"

// demo
type demo struct{}

func NewDemo() *demo {
	return &demo{}
}

func (a *demo) InitRouter(router *gin.Engine) {
	router.POST("/demo", a.demo)
}

func (a *demo) demo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hello xiet demo"})
}
