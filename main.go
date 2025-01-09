package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	////创建一个默认的路由引擎
	//r := gin.Default()
	//r.GET("/hello", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Hello  world",
	//	})
	//})
	//r.Run(":9090")
	core.InitializeViper()
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.Run(":9090")

}
