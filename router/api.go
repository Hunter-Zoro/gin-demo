package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/controller/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiGroupRouter(router *gin.RouterGroup) {
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.POST("/register", app.Register)
}
