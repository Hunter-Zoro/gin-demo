package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/controller/app"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
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
	router.POST("/login", app.Login)
	authRouter := router.Group("").Use(middleware.JWTAuth(service.AppGuardName))
	{
		authRouter.POST("/info", app.UserInfo)
	}
}
