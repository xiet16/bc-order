package handler

import "github.com/gin-gonic/gin"

// demo
type demo struct{}

func NewDemo() *demo {
	return &demo{}
}

func (a *demo) InitRouter(router *gin.Engine) {
	router.POST("/demo", a.demo)
	router.GET("/hello", a.hello)
}

func (a *demo) demo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hello xiet demo"})
}

func (a *demo) hello(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hello xiet"})
}
